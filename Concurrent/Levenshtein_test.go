package Concurrent

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

func TestMatchingData(t *testing.T) {
	MatchingData()
}

func TestMatchingDataParallel(t *testing.T) {
	MatchingDataParallel()
}
