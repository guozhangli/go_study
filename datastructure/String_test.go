package TestProject

import "testing"

func TestString_pattern(t *testing.T) {
	index := BruteForceStringMatch("goodgoogle", "google")
	t.Log(index)
}

func TestString_pattern2(t *testing.T) {
	var str MyString = "abaabbabaab"
	var F []int
	F = BuildFArray(str)
	t.Log(F)
}

func TestMyString_IndexKMP(t *testing.T) {
	var b MyString = "abaabaabbabaaabaabbabaab"
	var str MyString = "abaabbabaab"
	index := b.IndexKMP(str)
	t.Log(index)
}
