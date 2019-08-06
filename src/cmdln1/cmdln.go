package main

import ("os"
	"fmt"
)

func main() {
	argv := os.Args
	argc := len(argv)
	fmt.Printf("ArgC: %d(%d) Name: %s\n", argc, cap(argv),argv[0])
	for i,v := range argv {
		fmt.Printf("argv[%d] = %s (%T)\n", i, v, v)
	}
}
