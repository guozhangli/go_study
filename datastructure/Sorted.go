package TestProject

//冒泡排序
func BubblingSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
