package dataTest

import (
	"github.com/tadeuszjt/data"
	"testing"
)

func rowFloat32Identical(a, b data.RowT[float32]) bool {
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

func TestRowFloat32Identical(t *testing.T) {
	cases := []struct {
		a, b   data.RowT[float32]
		result bool
	}{
		{
			data.RowT[float32]{},
			data.RowT[float32]{},
			true,
		},
		{
			data.RowT[float32]{12},
			data.RowT[float32]{},
			false,
		},
		{
			data.RowT[float32]{1, 2, 3, 4},
			data.RowT[float32]{1, 2, 3, 4},
			true,
		},
		{
			data.RowT[float32]{1, 2, 3, 4},
			data.RowT[float32]{1, 2, 4, 4},
			false,
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := rowFloat32Identical(c.a, c.b)
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}

}

func TestRowFloat32Len(t *testing.T) {
	cases := []struct {
		row    data.RowT[float32]
		result int
	}{
		{data.RowT[float32]{}, 0},
		{data.RowT[float32]{1, 2, 3}, 3},
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

func TestRowFloat32Swap(t *testing.T) {
	cases := []struct {
		i, j   int
		row    data.RowT[float32]
		result data.RowT[float32]
	}{
		{0, 0, data.RowT[float32]{1}, data.RowT[float32]{1}},
		{0, 1, data.RowT[float32]{1, 2}, data.RowT[float32]{2, 1}},
		{1, 1, data.RowT[float32]{1, 2}, data.RowT[float32]{1, 2}},
		{1, 2, data.RowT[float32]{1, 2, 3, 4}, data.RowT[float32]{1, 3, 2, 4}},
		{3, 0, data.RowT[float32]{1, 2, 3, 4}, data.RowT[float32]{4, 2, 3, 1}},
	}

	for _, c := range cases {
		expected := c.result
		c.row.Swap(c.i, c.j)
		actual := c.row
		if !rowFloat32Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestRowFloat32Delete(t *testing.T) {
	cases := []struct {
		i      int
		row    data.RowT[float32]
		result data.RowT[float32]
	}{
		{0, data.RowT[float32]{1}, data.RowT[float32]{}},
		{1, data.RowT[float32]{1, 2, 3}, data.RowT[float32]{1, 3}},
		{1, data.RowT[float32]{1, 2, 3, 4}, data.RowT[float32]{1, 4, 3}},
		{3, data.RowT[float32]{1, 2, 3, 4}, data.RowT[float32]{1, 2, 3}},
		{0, data.RowT[float32]{1, 2, 3, 4}, data.RowT[float32]{4, 2, 3}},
	}

	for _, c := range cases {
		expected := c.result
		c.row.Delete(c.i)
		actual := c.row
		if !rowFloat32Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestRowFloat32Append(t *testing.T) {
	cases := []struct {
		f      float32
		row    data.RowT[float32]
		result data.RowT[float32]
	}{
		{0.1, data.RowT[float32]{}, data.RowT[float32]{0.1}},
		{0.2, data.RowT[float32]{1.0}, data.RowT[float32]{1.0, 0.2}},
		{0.3, data.RowT[float32]{1.0, 2.0, 3.0}, data.RowT[float32]{1.0, 2.0, 3.0, 0.3}},
	}

	for _, c := range cases {
		expected := c.result
		c.row.Append(c.f)
		actual := c.row
		if !rowFloat32Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}
