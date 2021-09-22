package bubble_sort

func Sort(arr []int64) []int64 {
	arrLen := len(arr)
	if arrLen < 1 {
		return arr
	}
	for i := arrLen - 1; i >= 0; i-- {
		flag := 0
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag++
			}
		}
		// 不存在交换，证明整个数组有序了
		if flag == 0 {
			break
		}
	}
	return arr
}
