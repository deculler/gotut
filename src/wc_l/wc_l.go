/*
  Package wc implements a simple record of word frequencies.

The database grows as words and their occurences are added.
This implementation illustrates the use of list container

*/

package wc_l

import (
	"fmt"
	"strings"
	"io"
	"container/list"
	"sort"
)

type WordCount struct {
	Word string
	Count int
}

func (wc *WordCount) AddCount(val int) {
	(*wc).Count +=  val
}

func (wc *WordCount) Inc() {
	wc.AddCount(1)
}

func (wc *WordCount) String() string {
	return fmt.Sprintf("%8d : %s", wc.Count, wc.Word)
}

type WordCounts struct {
	wcl list.List
}

// data interface
func (wcts *WordCounts) Len() int {
	return wcts.wcl.Len()
}

func (wcts *WordCounts) locate(i int) *list.Element {
	for j, e := 0, wcts.wcl.Front(); e != nil ; j, e = j+1, e.Next() {
		if (j == i) { 
			return e
		}
	}
	return nil 
}

func (wcts *WordCounts) Select(i int) *WordCount {
	return wcts.locate(i).Value.(*WordCount)
}

func (wcts *WordCounts) Less(i, j int) bool {
	wordi := wcts.Select(i).Word
	wordj := wcts.Select(j).Word
	return wordi < wordj
}

func (wcts *WordCounts) Swap(i, j int) {
	wci := wcts.locate(i)
	wcj := wcts.locate(j)
	wci.Value, wcj.Value = wcj.Value, wci.Value
}

func (wcts *WordCounts) Fprint(w io.Writer) {
	for e := wcts.wcl.Front(); e != nil; e = e.Next() {
		wc := e.Value.(*WordCount)
		fmt.Fprintf(w, "%8d: %s\n", wc.Count, wc.Word)
	}
}

func (wcts *WordCounts) Find(word string) *WordCount {
	for e := wcts.wcl.Front(); e != nil; e = e.Next() {
		wc := e.Value.(*WordCount)
		if strings.Compare(wc.Word, word) == 0 {
			return wc
		}
	}
	return nil
}

func (wcts *WordCounts) AddWord(word string, count int) {
	existing := wcts.Find(word)
	if existing == nil {
		wc := new(WordCount)
		wc.Word = word
		wc.Count = count
		wcts.wcl.PushFront(wc)
	} else {
		existing.AddCount(count)
	}
}


func (wcts *WordCounts) Sort() {
	sort.Sort(wcts)
}

func NewWordCounts() *WordCounts {
	/* Return an empty WordCounts */
	wct := new(WordCounts)
	wct.wcl.Init()
	fmt.Printf("Len %d\n", wct.Len())
	return wct
}
