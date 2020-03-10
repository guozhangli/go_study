package TestProject

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}
	for _, v := range arr {
		index := Binary_Search(arr, v)
		t.Log(index)
	}
}

func TestInterpolation_Search(t *testing.T) {
	arr := []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}
	for _, v := range arr {
		index := Interpolation_Search(arr, v)
		fmt.Println(index)
	}
}

func TestFabonacci_Search(t *testing.T) {
	arr := []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}
	for _, v := range arr {
		index := Fabonacci_Search(arr, v)
		fmt.Println(index)
	}
}
