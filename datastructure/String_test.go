package TestProject

import "testing"

func TestString_pattern(t *testing.T) {
	index := BruteForceStringMatch("goodgoogle", "google")
	t.Log(index)
}
