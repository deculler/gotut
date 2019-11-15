package wc_m

import (
	"os"
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	fmt.Printf("Testing wc_l.go\n")
	wct := NewWordCounts()	
	if wct.Len() != 0 {
		t.Error(`New WC not zero Len`)
 	}
	wct.Fprint(os.Stdout)
	wct.AddWord("Hello", 1)
	wct.Fprint(os.Stdout)
	wct.AddWord("World", 1)
	wct.Fprint(os.Stdout)
	wct.AddWord("Hello", 1)
	wct.Fprint(os.Stdout)
	wct.AddWord("Good", 1)
	if wct.Len() != 3 {
		t.Error(`New WC not length 3 after 3 inserts`)
 	}
	fmt.Println("Find")	
	fmt.Println(wct.Find("Hello"))
	wc := &WordCount{Word:"Hello", Count:2}
	if *wct.Find("Hello") != *wc {
		t.Error(`Find does not return correct WordCount`)
	}

	fmt.Println("Unsorted")
	wct.Fprint(os.Stdout)
	fmt.Println("Sorted")
	wct.Sort()
	wct.Fprint(os.Stdout)
}
