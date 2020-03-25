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

//归并排序
func MergeSort(arr []int) []int {
	var temp = make([]int, len(arr))
	MSort(arr, 0, len(arr)-1, temp)
	return temp
}

func MSort(arr []int, left int, right int, temp []int) {
	if left < right {
		mid := (left + right) / 2
		MSort(arr, left, mid, temp)
		MSort(arr, mid+1, right, temp)
		Marge(arr, left, mid, right, temp)
	}
}

func Marge(arr []int, left int, mid int, right int, temp []int) {
	i := left
	j := mid + 1
	t := 0
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[t] = arr[i]
			t++
			i++
		} else {
			temp[t] = arr[j]
			t++
			j++
		}
	}
	for i <= mid {
		temp[t] = arr[i]
		t++
		i++
	}
	for j <= right {
		temp[t] = arr[j]
		t++
		j++
	}
	t = 0
	for left <= right {
		arr[left] = temp[t]
		left++
		t++
	}
}

//快速排序
func QuickSort(arr []int, low int, high int) {
	if low < high {
		index := GetIndex(arr, low, high)
		QuickSort(arr, 0, index-1)
		QuickSort(arr, index+1, high)
	}
}

func GetIndex(arr []int, low int, high int) int {
	temp := arr[low]
	for low < high && temp <= arr[high] {
		high--
	}
	arr[low] = arr[high]
	for low < high && temp >= arr[low] {
		low++
	}
	arr[high] = arr[low]
	arr[low] = temp
	return low
}
