package TestProject

//折半查找
func Binary_Search(arr []int, key int) int {
	low := 0
	high := len(arr) - 1
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
	return -1
}

//插值查找
func Interpolation_Search(arr []int, key int) int {
	low := 0
	high := len(arr) - 1
	var mid int
	for low < high {
		mid = low + (key-arr[low])*(high-low)/(arr[high]-arr[low])
		if arr[mid] > key {
			high = mid - 1
		} else if arr[mid] < key {
			low = mid + 1
		} else if arr[mid] == key {
			return mid
		}
	}
	return -1
}

//斐波那契查找
func Fabonacci_Search(arr []int, key int) int {
	size := len(arr)
	fab := BuildFabArr(size - 1)
	fibBlock := len(fab) - 1
	left := 0
	right := size - 1
	var mid int
	for left <= right {
		mid = left + fab[fibBlock-1] - 1
		if key < arr[mid] {
			right = mid - 1
			fibBlock--
		} else if key > arr[mid] {
			left = mid + 1
			fibBlock -= 2
		} else {
			if mid >= size {
				return size - 1
			} else {
				return mid
			}
		}
	}
	return -1
}

func BuildFabArr(val int) []int {
	fab := make([]int, 0)
	fab = append(fab, 1)
	fab = append(fab, 1)
	i := 1
	for {
		fab = append(fab, fab[i]+fab[i-1])
		i++
		if fab[i] > val {
			break
		}
	}
	return fab
}
