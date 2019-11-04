package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"flag"
)

var wordflag bool


func rd_fun(infile *os.File) int {
	tokens := 0
	input := bufio.NewScanner(infile)
	if wordflag {
		input.Split(bufio.ScanWords)
	}
	for input.Scan() {
		fmt.Printf("%s\n", input.Text())
		tokens++
	}
	return tokens
}

func main() {
	flag.BoolVar(&wordflag, "words", false, "set output to words")
	flag.Parse()
	args := flag.Args()
	items := 0
	if len(args) == 0 {
		items += rd_fun(os.Stdin)
	} else {
		for _, name := range args {
			f, err := os.Open(name)
			if err != nil {
				log.Println("Bad filename ", name)
			}
			items += rd_fun(f)
			f.Close()
		}
	}
	if wordflag {
		fmt.Println(items, "words")
	} else {
		fmt.Println(items, "lines")
	}
}
