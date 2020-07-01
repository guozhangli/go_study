package geneticAlgorithm

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	strs := []string{
		"1", "2", "3", "4", "5", "6", "7", "8",
	}
	a, _ := ShuffleStrings(strs)
	fmt.Println(a)
}

func TestSerialGeneticAlgorithm(t *testing.T) {
	SerialGeneticAlgorithm()
}
