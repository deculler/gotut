package main

import ("os"
	"fmt"
	"reflect"
)

func main() {
	argv := os.Args
	argc := len(argv)
	fmt.Printf("ArgC: %d(%d) Name: %s\n", argc, cap(argv),argv[0])
	for i,v := range argv {
		fmt.Printf("argv[%d] = %s (%T)\n", i, v, v)
	}

	fmt.Printf("Args type: %T\n", argv)
	fmt.Println(argv)

	var pname [1]string = [...]string{argv[0] }
	fmt.Println(reflect.TypeOf(argv), reflect.TypeOf(pname))
	var chars = []byte(argv[0])
	fmt.Println(chars)
}
