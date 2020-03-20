package TestProject

import "fmt"

//冒泡排序
func BubblingSort(arr []int) []int {
	var flag = true
	for i := 0; (i < len(arr)) && flag; i++ {
		flag = false
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
				flag = true
			}
		}
	}
	return arr
}

//简单选择排序
func SimpleSearchSort(arr []int) []int {
	var min int
	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

//插入排序
func InsertSort(arr []int) []int {
	var temp int
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			temp = arr[i]
			var j int
			for j = i - 1; j >= 0 && arr[j] > temp; j-- {
				arr[j+1] = arr[j]
			}
			arr[j+1] = temp
		}
	}
	return arr
}

//希尔排序
func ShellSort(arr []int) []int {
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			InsertI(arr, gap, i)
		}
	}
	return arr
}
func InsertI(arr []int, gap int, i int) {
	var temp = arr[i]
	var j int
	for j = i - gap; j >= 0 && arr[j] > temp; j -= gap {
		arr[j+gap] = arr[j]
	}
	arr[j+gap] = temp
}

//堆排序
func HeapSort(arr []int) []int {
	n := len(arr)
	heap := NewHeap(n)
	heap.BuildHeap(arr)
	fmt.Println(heap.Array)
	for i := n - 1; i >= 0; i-- {
		heap.Array[0], heap.Array[i] = heap.Array[i], heap.Array[0]
		heap.Count--
		heap.PercolateDown(0)
	}
	return heap.Array
}
