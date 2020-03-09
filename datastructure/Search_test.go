package TestProject

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}
	index := Binary_Search(arr, 0)
	t.Log(index)
}

func TestInterpolation_Search(t *testing.T) {
	arr := []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}
	index := Interpolation_Search(arr, 62)
	t.Log(index)
}
