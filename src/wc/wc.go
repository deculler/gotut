/*
  Package wc implements a simple record of word frequencies.

The database grows as words and their occurences are added.
This implementation illustrates the use of slices.

*/

package wc

import (
	"fmt"
	"strings"
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

func (wc *WordCount) Stringer() string {
	return fmt.Sprintf("%8d : %s", wc.Count, wc.Word)
}

type WordCounts struct {
	wcs []WordCount
}

func (wcounts *WordCounts) Len() int {
	return len(wcounts.wcs)
}

func (wcounts *WordCounts) Print() {
	for _, wc := range wcounts.wcs {
		fmt.Println(wc)
	}
}

func (wcounts *WordCounts) Find(word string) *WordCount {
	for i := 0; i < len(wcounts.wcs); i++ {
		var wc *WordCount = &(wcounts.wcs[i])
		if strings.Compare(wc.Word, word) == 0 {
			return wc
		}
	}
	return nil
}

func (wcounts *WordCounts) AddWord(word string, count int) {
	fmt.Println("wc AddWord", wcounts.Len(), word, count)
	existing := wcounts.Find(word)
	if existing == nil {
		wc := WordCount{word, count}
		wcounts.wcs = append(wcounts.wcs, wc)
	} else {
		existing.AddCount(count)
	}
	wcounts.Print()
}

func NewWordCounts() WordCounts {
	/* Return an empty WordCounts */
	return WordCounts{[]WordCount{} }
}



