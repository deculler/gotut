/*
 Cmdln1 prints out the command line exec and args

*/

package main

import ("os"
	"fmt"
)

func main() {
	argv := os.Args    // Args is in the os package
	// argv is a slice of an array of strings.
	// They are typed, with length and capacity

	argc := len(argv)  // argc is not special 

	fmt.Printf("ArgC: %d, Cmd: %s, argv type: %T and capacity: %d\n", argc,  argv[0],
		argv, cap(argv))

	// idiomatic iteration in Go is to use range, either index or val can be _
	for i, v := range argv[1:] {
		fmt.Printf("argv[%d] = %s (%T)\n", i, v, v)
	}

	// the 3-expression for construct has ; but no enclosing ()
	for i := 0; i < argc; i++ {
		fmt.Printf("argv[%d] = %s\n", i, argv[i])
	}
}
