package levenshtein

import (
	"fmt"
	"testing"
)

/**
[0 1 2 3]
[1 0 1 2]
[2 1 0 1]
[3 2 1 1]
*/
func TestLevenshtein(t *testing.T) {
	v := Levenshtein("abc", "abe")
	fmt.Printf("%f\n", v)
}

/**
=== RUN   TestMatchingData
Dictionary Size:  250353
Word: stitter
Minimun distance: 1
List of best matching words: 9
sitter
skitter
slitter
spitter
stilter
stinter
stotter
stutter
titter
Execution Time: 194ms
--- PASS: TestMatchingData (0.23s)
PASS
*/
func TestMatchingData(t *testing.T) {
	MatchingData()
}

/**
=== RUN   TestMatchingDataParallel
Dictionary Size:  250353
2020/06/15 16:18:27 all goroutine execute finish
2020/06/15 16:18:27 worker num:0
Word: stitter
Minimun distance: 1
List of best matching words: 9
sitter
skitter
slitter
titter
stutter
stilter
stinter
stotter
spitter
Execution Time: 171ms
--- PASS: TestMatchingDataParallel (2.17s)
PASS
*/
func TestMatchingDataParallel(t *testing.T) {
	MatchingDataParallel()
}

/**
=== RUN   TestMatchingDataParallelFuture
Dictionary Size:  250353
2020/06/15 16:17:38 worker num:0
Word: stitter
Minimun distance: 1
List of best matching words: 9
sitter
skitter
slitter
spitter
stilter
stinter
stotter
stutter
titter
Execution Time: 149ms
2020/06/15 16:17:38 all goroutine execute finish
2020/06/15 16:17:38 worker num:0
--- PASS: TestMatchingDataParallelFuture (0.15s)
PASS
*/
func TestMatchingDataParallelFuture(t *testing.T) {
	MatchingDataParallelFuture()
}
