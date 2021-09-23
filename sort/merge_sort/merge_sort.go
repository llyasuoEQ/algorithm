package merge_sort

// Sort 归并排序
func Sort(arr []int64) []int64 {
	arrLen := len(arr)
	if arrLen < 2 {
		return arr
	}
	middle := arrLen / 2
	return merge(Sort(arr[:middle]), Sort(arr[middle:]))
}

// merge 合并数组
func merge(left, right []int64) []int64 {
	var result []int64
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}
