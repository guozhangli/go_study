package textindexing

import "testing"

/**
Execution Time: 425436
invertedIndex: 589624
goos: windows
goarch: amd64
pkg: go_study/Concurrent
BenchmarkTextIndexing-8   	       1	425535000000 ns/op
PASS
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
