package dataTest

import (
	"testing"
	"github.com/tadeuszjt/data"
)

func rowIntIdentical(a, b data.RowT[int]) bool {
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

func TestRowIntIdentical(t *testing.T) {
	cases := []struct{
		a, b   data.RowT[int]
		result bool
	}{
		{
			data.RowT[int]{},
			data.RowT[int]{},
			true,
		},
		{
			data.RowT[int]{12},
			data.RowT[int]{},
			false,
		},
		{
			data.RowT[int]{1, 2, 3, 4},
			data.RowT[int]{1, 2, 3, 4},
			true,
		},
		{
			data.RowT[int]{1, 2, 3, 4},
			data.RowT[int]{1, 2, 4, 4},
			false,
		},
	}
	
	for _, c := range cases {
		expected := c.result
		actual := rowIntIdentical(c.a, c.b)
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestRowIntLen(t *testing.T) {
	cases := []struct{
		row  data.RowT[int]
		result int
	}{
		{data.RowT[int]{}, 0},
		{data.RowT[int]{1, 2, 3}, 3},
		{nil, 0},
	}
	
	for _, c := range cases {
		expected := c.result
		actual := c.row.Len()
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestRowIntSwap(t *testing.T) {
	cases := []struct{
		i, j   int
		row  data.RowT[int]
		result data.RowT[int]
	}{
		{0, 0, data.RowT[int]{1}, data.RowT[int]{1}},
		{0, 1, data.RowT[int]{1, 2}, data.RowT[int]{2, 1}},
		{1, 1, data.RowT[int]{1, 2}, data.RowT[int]{1, 2}},
		{1, 2, data.RowT[int]{1, 2, 3, 4}, data.RowT[int]{1, 3, 2, 4}},
		{3, 0, data.RowT[int]{1, 2, 3, 4}, data.RowT[int]{4, 2, 3, 1}},
	}
	
	for _, c := range cases {
		expected := c.result
		c.row.Swap(c.i, c.j)
		actual := c.row
		if !rowIntIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestRowIntDelete(t *testing.T) {
	cases := []struct{
		i      int
		row  data.RowT[int]
		result data.RowT[int]
	}{
		{0, data.RowT[int]{1}, data.RowT[int]{}},
		{1, data.RowT[int]{1, 2, 3}, data.RowT[int]{1, 3}},
		{1, data.RowT[int]{1, 2, 3, 4}, data.RowT[int]{1, 4, 3}},
		{3, data.RowT[int]{1, 2, 3, 4}, data.RowT[int]{1, 2, 3}},
		{0, data.RowT[int]{1, 2, 3, 4}, data.RowT[int]{4, 2, 3}},
	}
	
	for _, c := range cases {
		expected := c.result
		c.row.Delete(c.i)
		actual := c.row
		if !rowIntIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}
