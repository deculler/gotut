package wc_l

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
	wct.Fprint(os.Stdout)
	if wct.Len() != 3 {
		t.Error(`New WC not length 3 after 3 inserts`)
 	}
	fmt.Println(wct.Select(0), wct.Select(1), wct.Select(2))
	fmt.Println(wct.Less(0,1))
	wct.Swap(0,1)
	fmt.Println(wct.Select(0), wct.Select(1), wct.Select(2))
	wct.Sort()
	wct.Fprint(os.Stdout)
	if !wct.Less(0,2) {
		t.Error(`Not sorted`)
	}
}
