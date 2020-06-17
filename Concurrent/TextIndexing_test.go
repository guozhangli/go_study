package Concurrent

import "testing"

/**
=== RUN   TestTextIndexing
Execution Time: 966414
invertedIndex: 647673
--- PASS: TestTextIndexing (966.51s)
*/
func BenchmarkTextIndexing(b *testing.B) {
	TextIndexingSerial()
}

func TestTextIndexingParallel(t *testing.T) {
	TextIndexingParallel()
}

func TestTextIndexingGroup(t *testing.T) {
	TextIndexingGroup()
}
