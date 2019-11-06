/*

  wordct command, counting the frequency of words in a collection of files
using a go routine per file to perform the parsing and a go routine to 
accumulate all the counts

*/

package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	wc "wc_s"
	//wc "wc_l"
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

func readwords(infile *os.File, name string, ch_adder chan string, done chan string) {
	fmt.Println("Go: ", name)
	rdr := bufio.NewReader(infile)    // bufio, but Reader interface
	for wrd := getword(rdr); len(wrd) > 0; wrd = getword(rdr) {
		str := strings.ToLower(string(wrd))
		if (len(str) > 1) {
			ch_adder <- str
		}
	}
	infile.Close()
	done <- name

}

func adder(wcounts *wc.WordCounts,  ch_adder chan string, done chan string) {
	fmt.Println("Go adder")
	for str := range ch_adder { // accumulate all the parsed words
		wcounts.AddWord(str, 1)
	}
	done <- "done"	
}

func main() {
	args := os.Args
	wordcounts := wc.NewWordCounts()
	if len(args) == 1 {
		countwords(os.Stdin, wordcounts)
	} else {
		addr := make(chan string)
		rdone := make(chan string) // reader/producers done
		adone := make(chan string) // consumer done

		for _, name := range args[1:] {
			f, err := os.Open(name)
			if err != nil {
				log.Println("Bad filename ", name)
			} else {
				go readwords(f, name, addr, rdone)
			}
		}
		go adder(wordcounts, addr, adone)

		// Receive done from all the readers
		for i, _ := range args[1:] {
			n := <- rdone
			fmt.Println("Done: ", i, n)
		}
		close(addr)	// close the channel, terminating adder
		n := <- adone	// Receive done from adder
		fmt.Println("adder", n)
	}
	wordcounts.Sort()
	wordcounts.Fprint(os.Stdout)
}
