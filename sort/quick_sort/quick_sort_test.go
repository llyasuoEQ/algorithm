package quick_sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	arr := []int64{2, 1, 3, 6}
	actual := Sort(arr)
	expect := []int64{1, 2, 3, 6}
	assert.Equal(t, expect, actual, "Sort execute error")
}

func BenchmarkSort(b *testing.B) {
	arr := []int64{2, 1, 3, 6, 2, 3}
	for i := 0; i < b.N; i++ {
		tmp := make([]int64, len(arr))
		copy(tmp, arr)
		Sort(tmp)
	}
}
