/*

  wordct command, counting the frequency of words in a collection of files
using a go routine per file to perform the parsing and update a shared
word-count data structure

*/

package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	wc "wc_sm"
)

func isalpha(c byte) bool {
	return ((c >= 'a') && (c <= 'z') || ((c >= 'A') && (c <= 'Z')))
}

func getword(rdr *bufio.Reader) string {
	/*
	  Stateless tokenizer - next alpha string, skipping non-alpha
	*/
	var ch byte
	var err error
	word := ""		// string can be dynamically allocated
	ch, err = rdr.ReadByte()
	for err == nil && !isalpha(ch) {
		ch, err = rdr.ReadByte() // Alternative is ReadRune 
	}

	for  err == nil && isalpha(ch) {
		word += string(ch) // grow via append
		ch, err = rdr.ReadByte()		
	}
	return word		
}


func countwords(infile *os.File, name string,  wcounts *wc.WordCounts, done chan string) {
	/* Return slice of unique words in a File. */
	rdr := bufio.NewReader(infile)    // bufio, but Reader interface
	for wrd := getword(rdr); len(wrd) > 0; wrd = getword(rdr) {
		str := strings.ToLower(string(wrd))
		if (len(str) > 1) {
			wcounts.AddWord(str, 1)
		}
	}
	if name != "" {
		infile.Close()
		done <- name
	}
}


func main() {
	args := os.Args
	wordcounts := wc.NewWordCounts()
	rdone := make(chan string) // reader/producers done
	if len(args) == 1 {
		countwords(os.Stdin, "", wordcounts, rdone)
	} else {
		for _, name := range args[1:] {
			f, err := os.Open(name)
			if err != nil {
				log.Println("Bad filename ", name)
			} else {
				go countwords(f, name, wordcounts, rdone)
			}
		}

		// Receive done from all the readers
		for i, _ := range args[1:] {
			n := <- rdone
			fmt.Println("Done: ", i, n)
		}
	}
	wordcounts.Sort()
	wordcounts.Fprint(os.Stdout)
}
