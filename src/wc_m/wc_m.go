/*
  Package wc implements a simple record of word frequencies.

The database grows as words and their occurences are added.
This implementation illustrates the use of maps

*/

package wc_m

import (
	"sort"
	"fmt"
	"io"
)

// WordCount illustrates the basic definition of a type and set of methods on the type

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

/* WordCounts illustrates thge use of a map to form the set of word:counts 
    - note that the WordCount type does not appear in WordCounts, as it would for 
slice or list based representations
*/

type WordCounts struct {
	wcs map[string]int
	srt bool
}

func (wcts *WordCounts) Len() int {
	return len(wcts.wcs)
}

func (wcts *WordCounts) Fprint(w io.Writer) {
	if wcts.srt {
		swc := make([]WordCount, 0, wcts.Len())
		for word, count := range wcts.wcs {
			swc = append(swc, WordCount{Word:word, Count:count})
		}
		wc_less := func(i, j int) bool{
			return swc[i].Word < swc[j].Word
		}
		sort.Slice(swc, wc_less)
		for _, wc := range swc {
			fmt.Fprintf(w, "%8d: %s\n", wc.Count, wc.Word)
		}
		
	} else {
		for word, count := range wcts.wcs {
			fmt.Fprintf(w, "%8d: %s\n", count, word)
		}
	}
}

func (wcts *WordCounts) Find(word string) *WordCount {
	// Return a reference to a WordCount for word, if it exists
	ct, ok := wcts.wcs[word]
	if ok {
		return &WordCount{Word : word, Count : ct}
	} else {
		// If the key is not present, indexing wiht it returns the zero of the val type
		return nil
	}
	
}

func (wcts *WordCounts) AddWord(word string, count int) {
	ct, ok := wcts.wcs[word]
	if ok {
		wcts.wcs[word] = ct + count
	} else {
		wcts.wcs[word] = count
	}
}

func (wcts *WordCounts) Sort() {
	wcts.srt = true
}

func NewWordCounts() *WordCounts {
	/* Return an empty WordCounts */
	wct := new(WordCounts)
	wct.srt = false
	wct.wcs = make(map[string]int, 100)
	return wct
}



