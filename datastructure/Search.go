package TestProject

//折半查找
func Binary_Search(arr []int, key int) int {
	low := 1
	high := len(arr)
	var mid int
	for low < high {
		mid = (low + high) / 2
		if arr[mid] > key {
			high = mid
		} else if arr[mid] < key {
			low = mid
		} else if arr[mid] == key {
			return mid
		}
	}
	return 0
}

//插值查找
func Interpolation_Search(arr []int, key int) int {
	low := 1
	high := len(arr) - 1
	var mid int
	for low < high {
		mid = low + (key-arr[low])*(high-low)/(arr[high]-arr[low])
		if arr[mid] > key {
			high = mid
		} else if arr[mid] < key {
			low = mid
		} else if arr[mid] == key {
			return mid
		}
	}
	return 0
}
