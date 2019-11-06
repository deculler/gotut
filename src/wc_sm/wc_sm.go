/*
  Package wc implements a simple record of word frequencies.

The database grows as words and their occurences are added.
This implementation illustrates the use of slices.

*/

package wc_sm

import (
	"fmt"
	"strings"
	"io"
	"sort"
	"sync"
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
	mu sync.Mutex
	wcs []WordCount
}

// data interface - not thread safe
func (wcts *WordCounts) Len() int {
	return len(wcts.wcs)
}

func (wcts *WordCounts) Less(i, j int) bool {
	return wcts.wcs[i].Word < wcts.wcs[j].Word
}

func (wcts *WordCounts) Swap(i, j int) {
	wcts.wcs[i], wcts.wcs[j] = wcts.wcs[j], wcts.wcs[i]
}

func (wcts *WordCounts) Fprint(w io.Writer) {
	for _, wc := range wcts.wcs {
		fmt.Fprintf(w, "%8d: %s\n", wc.Count, wc.Word)
	}
}

func (wcts *WordCounts) Find(word string) *WordCount {
	for i := 0; i < len(wcts.wcs); i++ {
		var wc *WordCount = &(wcts.wcs[i])
		if strings.Compare(wc.Word, word) == 0 {
			return wc
		}
	}
	return nil
}

func (wcts *WordCounts) AddWord(word string, count int) {
	wcts.mu.Lock()
	existing := wcts.Find(word)
	if existing == nil {
		wc := WordCount{word, count}
		wcts.wcs = append(wcts.wcs, wc)
	} else {
		existing.AddCount(count)
	}
	wcts.mu.Unlock()
}

func (wcts *WordCounts) Sort() {
	sort.Sort(wcts)
}

func NewWordCounts() *WordCounts {
	/* Return an empty WordCounts */
	return &WordCounts{[]WordCount{} }
}



