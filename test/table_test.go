package dataTest

import (
	. "github.com/tadeuszjt/data"
	"testing"
)

func tableIdentical(a, b Table) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		switch sa := a[i].(type) {
		case *RowT[int]:
			sb, ok := b[i].(*RowT[int])
			if !ok || !rowIntIdentical(*sa, *sb) {
				return false
			}

		case *RowT[float32]:
			sb, ok := b[i].(*RowT[float32])
			if !ok || !rowFloat32Identical(*sa, *sb) {
				return false
			}

		default:
			panic("testRowIntIdentical: unrecognised row type")
		}
	}

	return true
}

func TestTableIdentical(t *testing.T) {
	cases := []struct {
		a, b   Table
		result bool
	}{
		{
			Table{},
			Table{},
			true,
		},
		{
			Table{},
			Table{&RowT[int]{}},
			false,
		},
		{
			Table{&RowT[int]{}},
			Table{&RowT[int]{}},
			true,
		},
		{
			Table{&RowT[float32]{}},
			Table{&RowT[int]{}},
			false,
		},
		{
			Table{&RowT[int]{1, 2, 3}},
			Table{&RowT[int]{1, 2, 3}},
			true,
		},
		{
			Table{&RowT[int]{1, 2, 3}},
			Table{&RowT[int]{1, 2, 4}},
			false,
		},
		{
			Table{
				&RowT[int]{1, 2, 3},
				&RowT[float32]{1, 2, 3},
			},
			Table{
				&RowT[int]{1, 2, 3},
				&RowT[float32]{1, 2, 3},
			},
			true,
		},
		{
			Table{
				&RowT[int]{1, 2, 3},
				&RowT[float32]{1, 2, 3},
			},
			Table{
				&RowT[int]{1, 2, 3},
				&RowT[float32]{1, 2, 3.1},
			},
			false,
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := tableIdentical(c.a, c.b)
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestTableLen(t *testing.T) {
	cases := []struct {
		table  Table
		result int
	}{
		{
			Table{&RowT[int]{}},
			0,
		},
		{
			Table{&RowT[int]{1, 2, 3}},
			3,
		},
		{
			Table{
				&RowT[int]{1, 2, 3},
				&RowT[float32]{1, 2, 3},
			},
			3,
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.table.Len()
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestTableSwap(t *testing.T) {
	cases := []struct {
		i, j   int
		table  Table
		result Table
	}{
		{
			0, 0,
			Table{},
			Table{},
		},
		{
			0, 0,
			Table{&RowT[int]{1}},
			Table{&RowT[int]{1}},
		},
		{
			1, 3,
			Table{&RowT[int]{1, 2, 3, 4}},
			Table{&RowT[int]{1, 4, 3, 2}},
		},
		{
			2, 0,
			Table{
				&RowT[int]{1, 2, 3, 4},
				&RowT[float32]{.1, .2, .3, .4},
			},
			Table{
				&RowT[int]{3, 2, 1, 4},
				&RowT[float32]{.3, .2, .1, .4},
			},
		},
	}

	for _, c := range cases {
		expected := c.result
		c.table.Swap(c.i, c.j)
		actual := c.table
		if !tableIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestTableDelete(t *testing.T) {
	cases := []struct {
		i      int
		table  Table
		result Table
	}{
		{
			0,
			Table{},
			Table{},
		},
		{
			0,
			Table{
				&RowT[int]{1, 2, 3, 4},
				&RowT[float32]{1, 2, 3, 4},
			},
			Table{
				&RowT[int]{4, 2, 3},
				&RowT[float32]{4, 2, 3},
			},
		},
		{
			1,
			Table{
				&RowT[int]{1, 2, 3, 4},
				&RowT[float32]{1, 2, 3, 4},
			},
			Table{
				&RowT[int]{1, 4, 3},
				&RowT[float32]{1, 4, 3},
			},
		},
		{
			3,
			Table{
				&RowT[int]{1, 2, 3, 4},
				&RowT[float32]{1, 2, 3, 4},
			},
			Table{
				&RowT[int]{1, 2, 3},
				&RowT[float32]{1, 2, 3},
			},
		},
	}

	for _, c := range cases {
		expected := c.result
		c.table.Delete(c.i)
		actual := c.table
		if !tableIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}
