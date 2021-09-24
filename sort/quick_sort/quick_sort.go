package quick_sort

// Sort 快速排序
func Sort(arr []int64) []int64 {
	arrLen := len(arr)
	if arrLen < 1 {
		return arr
	}
	return quick(arr, 0, arrLen-1)
}

func quick(arr []int64, left, right int) []int64 {
	if left < right {
		// 寻找中间值下标
		arr, mid := patition(arr, left, right)
		// 左数组进行快排
		quick(arr, left, mid-1)
		// 右数组进行快排
		quick(arr, mid+1, right)
	}
	return arr
}

func patition(arr []int64, left, right int) ([]int64, int) {
	p := arr[left] // 基准
	i, j := left, right
	for i < j {
		// 以最左边的值为基准，那么先从最有边开始找小于基准的
		// 如果以左右边值为基准，那么先从最左边开始找大于基准的
		for arr[j] >= p && i < j {
			j--
		}
		for arr[i] <= p && i < j {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[left], arr[i] = arr[i], arr[left]
	return arr, i
}
