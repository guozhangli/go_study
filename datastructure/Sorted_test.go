package TestProject

import "testing"

func TestBubblingSorted(t *testing.T) {
	var arr = []int{2, 4, 5, 9, 3, 6, 7, 1, 10, 8}
	arrsorted := BubblingSort(arr)
	t.Log(arrsorted)
}

func TestSimpleSearchSorted(t *testing.T) {
	var arr = []int{2, 4, 5, 9, 3, 6, 7, 1, 10, 8}
	arrsorted := SimpleSearchSort(arr)
	t.Log(arrsorted)
}

func TestInsertSorted(t *testing.T) {
	var arr = []int{2, 4, 5, 9, 3, 6, 7, 1, 10, 8}
	arrsorted := InsertSort(arr)
	t.Log(arrsorted)
}

func TestShellSorted(t *testing.T) {
	var arr = []int{2, 4, 5, 9, 3, 6, 7, 1, 10, 8}
	arrsorted := ShellSort(arr)
	t.Log(arrsorted)
}

func TestHeapSorted(t *testing.T) {
	var arr = []int{2, 4, 5, 9, 3, 6, 7, 1, 10, 8}
	arrsorted := HeapSort(arr)
	t.Log(arrsorted)
}

func TestMergeSorted(t *testing.T) {
	var arr = []int{2, 4, 5, 9, 3, 6, 7, 1, 10, 8}
	arrsorted := MergeSort(arr)
	t.Log(arrsorted)
}

func TestQuickSorted(t *testing.T) {
	var arr = []int{2, 4, 5, 9, 3, 6, 7, 1, 10, 8}
	QuickSort(arr, 0, len(arr)-1)
	t.Log(arr)
}
