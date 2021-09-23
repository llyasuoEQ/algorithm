package shell_sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	arr1 := []int64{2, 1, 3, 6, 2, 3}
	arr2 := []int64{2, 1, 3, 6, 2, 3}
	actual1 := Sort(arr1)
	actual2 := HibbardSort(arr2)
	expect := []int64{1, 2, 2, 3, 3, 6}
	assert.Equal(t, expect, actual1, "Sort execute error")
	assert.Equal(t, expect, actual2, "HibbardSort execute error")
}

func BenchmarkSort(b *testing.B) {
	arr := []int64{2, 1, 3, 6, 2, 3}

	for i := 0; i < b.N; i++ {
		tmp := []int64{}
		for _, v := range arr {
			tmp = append(tmp, v)
		}
		Sort(tmp)
	}

}

func BenchmarkHibbardSort(b *testing.B) {
	arr := []int64{2, 1, 3, 6, 2, 3}
	for i := 0; i < b.N; i++ {
		tmp := make([]int64, len(arr))
		copy(tmp, arr)
		HibbardSort(tmp)
	}
}
