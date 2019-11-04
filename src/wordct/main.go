package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"fmt"
	"wc"
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
	// no need to have provided buffer for string object in outer scope or malloc
	return word		
}

func countwords(infile *os.File, wcounts *wc.WordCounts) {
	/* Return slice of unique words in a File. */
	rdr := bufio.NewReader(infile)    // bufio, but Reader interface
	for wrd := getword(rdr); len(wrd) > 0; wrd = getword(rdr) {
		str := strings.ToLower(string(wrd))
		if (len(str) > 1) {
			wcounts.AddWord(str, 1)
		}
	}
}

func main() {
	args := os.Args
	wordcounts := wc.NewWordCounts()
	if len(args) == 1 {
		countwords(os.Stdin, &wordcounts)
	} else {
		for _, name := range args[1:] {
			f, err := os.Open(name)
			if err != nil {
				log.Println("Bad filename ", name)
			} else {
				countwords(f, &wordcounts)
				f.Close()
			}
		}
	}
	fmt.Println("main done")
	wordcounts.Print()
}
