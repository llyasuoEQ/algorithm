package insert_sort

// Sort 冒泡排序
func Sort(arr []int64) []int64 {
	arrLen := len(arr)
	if arrLen < 1 {
		return arr
	}
	for i := 1; i < arrLen; i++ {
		tmp := arr[i]
		j := i
		for ; (j > 0) && (tmp < arr[j-1]); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
	return arr
}
