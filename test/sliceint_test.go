package dataTest

import (
	"testing"
	"github.com/tadeuszjt/data"
)

func sliceIntIdentical(a, b data.SliceT[int]) bool {
	if len(a) != len(b) {
		return false
	}
	
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	
	return true
}

func TestSliceIntIdentical(t *testing.T) {
	cases := []struct{
		a, b   data.SliceT[int]
		result bool
	}{
		{
			data.SliceT[int]{},
			data.SliceT[int]{},
			true,
		},
		{
			data.SliceT[int]{12},
			data.SliceT[int]{},
			false,
		},
		{
			data.SliceT[int]{1, 2, 3, 4},
			data.SliceT[int]{1, 2, 3, 4},
			true,
		},
		{
			data.SliceT[int]{1, 2, 3, 4},
			data.SliceT[int]{1, 2, 4, 4},
			false,
		},
	}
	
	for _, c := range cases {
		expected := c.result
		actual := sliceIntIdentical(c.a, c.b)
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestSliceIntLen(t *testing.T) {
	cases := []struct{
		slice  data.SliceT[int]
		result int
	}{
		{data.SliceT[int]{}, 0},
		{data.SliceT[int]{1, 2, 3}, 3},
		{nil, 0},
	}
	
	for _, c := range cases {
		expected := c.result
		actual := c.slice.Len()
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestSliceIntSwap(t *testing.T) {
	cases := []struct{
		i, j   int
		slice  data.SliceT[int]
		result data.SliceT[int]
	}{
		{0, 0, data.SliceT[int]{1}, data.SliceT[int]{1}},
		{0, 1, data.SliceT[int]{1, 2}, data.SliceT[int]{2, 1}},
		{1, 1, data.SliceT[int]{1, 2}, data.SliceT[int]{1, 2}},
		{1, 2, data.SliceT[int]{1, 2, 3, 4}, data.SliceT[int]{1, 3, 2, 4}},
		{3, 0, data.SliceT[int]{1, 2, 3, 4}, data.SliceT[int]{4, 2, 3, 1}},
	}
	
	for _, c := range cases {
		expected := c.result
		c.slice.Swap(c.i, c.j)
		actual := c.slice
		if !sliceIntIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestSliceIntDelete(t *testing.T) {
	cases := []struct{
		i      int
		slice  data.SliceT[int]
		result data.SliceT[int]
	}{
		{0, data.SliceT[int]{1}, data.SliceT[int]{}},
		{1, data.SliceT[int]{1, 2, 3}, data.SliceT[int]{1, 3}},
		{1, data.SliceT[int]{1, 2, 3, 4}, data.SliceT[int]{1, 4, 3}},
		{3, data.SliceT[int]{1, 2, 3, 4}, data.SliceT[int]{1, 2, 3}},
		{0, data.SliceT[int]{1, 2, 3, 4}, data.SliceT[int]{4, 2, 3}},
	}
	
	for _, c := range cases {
		expected := c.result
		c.slice.Delete(c.i)
		actual := c.slice
		if !sliceIntIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}
