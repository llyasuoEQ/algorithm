package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBsSearch(t *testing.T) {
	arr, f := []int64{1, 2, 3, 4, 5, 6, 7}, int64(1)
	actual := BsSearch(arr, f)
	expect := 0
	assert.Equal(t, expect, actual, "BsSearch execute failed")
}

type IntSlice []int64

func (i IntSlice) len() int {
	return len(i)
}

func (i IntSlice) less(index int, f interface{}) bool {
	fi, ok := f.(int64)
	if !ok {
		return false
	}
	return i[index] < fi
}

func (i IntSlice) more(index int, f interface{}) bool {
	fi, ok := f.(int64)
	if !ok {
		return false
	}
	return i[index] > fi
}

func TestBSIterfaceSearch(t *testing.T) {
	arr, f := IntSlice{1, 2, 3, 4, 5, 6, 7}, int64(1)
	actual := BSIterfaceSearch(arr, f)
	expect := 0
	assert.Equal(t, expect, actual, "TestBSIterfaceSearch execute failed")
}
