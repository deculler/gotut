package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

func words(infile *os.File) []string {
	/* Return slice of unique words in a File. */
	rdr := bufio.NewReader(infile)    // bufio, but Reader interface
	uwords := make([]string, 0, 1024) // slice backed by dynamically allocated array
	for wrd := getword(rdr); len(wrd) > 0; wrd = getword(rdr) {
		str := strings.ToLower(string(wrd))
		if (len(str) > 1) {
			uwords = append(uwords, str) // grow as needed
		}
	}
	return uwords
}

func printwords(uniques []string) {
	for _, wrd := range uniques {
		fmt.Println(wrd)
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		printwords(words(os.Stdin))
	} else {
		for _, name := range args[1:] {
			f, err := os.Open(name)
			if err != nil {
				log.Println("Bad filename ", name)
			} else {
				printwords(words(f))
				f.Close()
			}
		}
	}
}
