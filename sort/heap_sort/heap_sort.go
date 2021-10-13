package heap_sort

// Sort 堆排序，升序
func Sort(arr []int64) []int64 {
	length := len(arr)
	if length < 2 {
		return arr
	}
	buildMaxHeap(arr, length) // 建一个大顶堆
	for i := length - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		adjustMaxHeap(arr, 0, i)
	}

	return arr
}

// buildMaxHeap 建堆
func buildMaxHeap(arr []int64, length int) {
	// 从最底层开始调整，然后依次向上，直到根节点
	for i := length / 2; i >= 0; i-- {
		// 调整 i 下标对应的左子树和右子树
		adjustMaxHeap(arr, i, length)
	}
}

// adjustMaxHeap 调整堆
func adjustMaxHeap(arr []int64, i, length int) {
	left := i*2 + 1
	right := i*2 + 2
	maxIndex := i                                   // 标记最大值节点
	if left < length && arr[left] > arr[maxIndex] { // 先对比左节点
		maxIndex = left
	}
	if right < length && arr[right] > arr[maxIndex] { // 对比右节点
		maxIndex = right
	}
	if maxIndex != i { // 最大值的下标变更，说明需要调整
		arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		// 在调整子节点的树，递归下去，需要一直调整，直到满足大顶堆的性质即可
		adjustMaxHeap(arr, maxIndex, length)
	}
}
