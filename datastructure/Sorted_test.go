package TestProject

import "testing"

func TestSortedBubbling(t *testing.T) {
	var arr = []int{2, 4, 5, 9, 3, 6, 7, 1, 10, 8}
	arrsorted := BubblingSort(arr)
	t.Log(arrsorted)
}
