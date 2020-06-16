package Concurrent

import "testing"

func TestTextIndexing(t *testing.T) {
	TextIndexingSerial()
}

func TestTextIndexingParallel(t *testing.T) {
	TextIndexingParallel()
}

func TestTextIndexingGroup(t *testing.T) {
	TextIndexingGroup()
}
