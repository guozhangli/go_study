package TestProject

import (
	"testing"
)

func TestNewHeap(t *testing.T) {
	heap := NewHeap(5)
	t.Log(heap)
}

func TestParentIndex(t *testing.T) {
	heap := NewHeap(5)
	index := heap.ParentIndex(3)
	t.Log(index)
}

func TestLeftChildIndex(t *testing.T) {
	heap := NewHeap(5)
	index := heap.LeftChildIndex(3)
	t.Log(index)
}

func TestRightChildIndex(t *testing.T) {
	heap := NewHeap(5)
	index := heap.RightChildIndex(3)
	t.Log(index)
}

func TestGetMaximum(t *testing.T) {
	heap := NewHeap(5)
	val := heap.GetMaximum()
	t.Log(val)
}

func TestPercolateDown(t *testing.T) {
	heap := NewHeap(5)
	heap.PercolateDown(3)
	t.Log(heap)
}

func TestDeleteMax(t *testing.T) {
	heap := NewHeap(5)
	val := heap.DeleteMax()
	t.Log(val)
}

func TestInsert(t *testing.T) {
	heap := NewHeap(5)
	heap.Insert(6)
	t.Log(heap)
}

func TestDestoryHeap(t *testing.T) {
	heap := NewHeap(5)
	heap.DestoryHeap()
	t.Log(heap)
}

func TestBuildHeap(t *testing.T) {
	var arr = []int{2, 4, 5, 9, 3, 6, 7, 1, 10, 8}
	heap := NewHeap(len(arr))
	heap.BuildHeap(arr)
	t.Log(heap)
}

func TestHeapSort(t *testing.T) {
	var arr = []int{2, 4, 5, 9, 3, 6, 7, 1, 10, 8}
	a := HeapSort(arr)
	t.Log(a)
}
