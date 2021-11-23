package binary_search

// bsSearch 二分查找
func BsSearch(arr []int64, f int64) int {
	res := -1
	length := len(arr)
	if len(arr) < 1 {
		return res
	}
	left := 0
	right := length - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > f {
			right = mid - 1
		} else if arr[mid] < f {
			left = mid + 1
		} else {
			res = mid
			break
		}
	}
	return res
}

// BSIterface 可以任意类型去用的函数底层函数
type BSIterface interface {
	len() int                   // 长度
	less(int, interface{}) bool // 小于
	more(int, interface{}) bool // 大于
}

// BSIterfaceSearch 二分查找
func BSIterfaceSearch(data BSIterface, f interface{}) int {
	res := -1
	length := data.len()
	if length < 1 {
		return res
	}
	left, right := 0, length-1
	for left <= right {
		mid := (left + right) / 2
		if data.less(mid, f) {
			left = mid + 1
		} else if data.more(mid, f) {
			right = mid - 1
		} else {
			res = mid
			break
		}
	}
	return res
}
