package shell_sort

// Sort 希尔排序
func Sort(arr []int64) []int64 {
	arrLen := len(arr)
	if arrLen < 1 {
		return arr
	}
	for i := arrLen / 2; i > 0; i = i / 2 {
		// 插入排序
		for k := i; k < arrLen; k++ {
			tmp := arr[k]
			m := k
			for ; (m >= i) && (arr[m-i] > tmp); m-- {
				arr[m] = arr[m-i]
			}
			arr[m] = tmp
		}
	}
	return arr
}

// Hibbard 算法
// Sort 希尔排序
func HibbardSort(arr []int64) []int64 {
	arrLen := len(arr)
	if arrLen < 1 {
		return arr
	}
	gap := 1
	for ; gap < arrLen; gap = gap*2 + 1 {
	}
	for ; gap > 0; gap = (gap - 1) / 2 {
		for i := gap; i < arrLen; i++ {
			tmp := arr[i]
			j := i
			for ; j >= gap && arr[j-gap] > tmp; j-- {
				arr[j] = arr[j-gap]
			}
			arr[j] = tmp
		}
	}
	return arr
}
