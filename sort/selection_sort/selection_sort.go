package selection_sort

//sort 选择排序
func Sort(arr []int64) []int64 {
	arrLen := len(arr)
	if arrLen < 1 {
		return arr
	}
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			if arr[i] > arr[j] {
				// 对i,j位置的元素进行交换
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
