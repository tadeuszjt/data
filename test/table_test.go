package dataTest

import (
	"testing"
	"github.com/tadeuszjt/data"
)

func tableIdentical(a, b data.Table) bool {
	if len(a) != len(b) {
		return false
	}
	
	for i := range a {
		switch sa := a[i].(type) {
			case *data.SliceInt:
				sb, ok := b[i].(*data.SliceInt)
				if !ok || !sliceIntIdentical(*sa, *sb) {
					return false
				}
				
			case *data.SliceFloat32:
				sb, ok := b[i].(*data.SliceFloat32)
				if !ok || !sliceFloat32Identical(*sa, *sb) {
					return false
				}
				
			default:
				panic("testSliceIntIdentical: unrecognised slice type")
		}
	}
	
	return true
}

func TestTableIdentical(t *testing.T) {
	cases := []struct{
		a, b   data.Table
		result bool
	}{
		{
			data.Table{},
			data.Table{},
			true,
		},
		{
			data.Table{},
			data.Table{ &data.SliceInt{} },
			false,
		},
		{
			data.Table{ &data.SliceInt{} },
			data.Table{ &data.SliceInt{} },
			true,
		},
		{
			data.Table{ &data.SliceFloat32{} },
			data.Table{ &data.SliceInt{} },
			false,
		},
		{
			data.Table{ &data.SliceInt{1, 2, 3} },
			data.Table{ &data.SliceInt{1, 2, 3} },
			true,
		},
		{
			data.Table{ &data.SliceInt{1, 2, 3} },
			data.Table{ &data.SliceInt{1, 2, 4} },
			false,
		},
		{
			data.Table{
				&data.SliceInt{1, 2, 3},
				&data.SliceFloat32{1, 2, 3},
			},
			data.Table{
				&data.SliceInt{1, 2, 3},
				&data.SliceFloat32{1, 2, 3},
			},
			true,
		},
		{
			data.Table{
				&data.SliceInt{1, 2, 3},
				&data.SliceFloat32{1, 2, 3},
			},
			data.Table{
				&data.SliceInt{1, 2, 3},
				&data.SliceFloat32{1, 2, 3.1},
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
	cases := []struct{
		table  data.Table
		result int
	}{
		{
			data.Table{ &data.SliceInt{} },
			0,
		},
		{
			data.Table{ &data.SliceInt{1, 2, 3} },
			3,
		},
		{
			data.Table{
				&data.SliceInt{1, 2, 3},
				&data.SliceFloat32{1, 2, 3},
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
	cases := []struct{
		i, j   int
		table  data.Table
		result data.Table
	}{
		{
			0, 0,
			data.Table{},
			data.Table{},
		},
		{
			0, 0,
			data.Table{ &data.SliceInt{1} },
			data.Table{ &data.SliceInt{1} },
		},
		{
			1, 3,
			data.Table{ &data.SliceInt{1, 2, 3, 4} },
			data.Table{ &data.SliceInt{1, 4, 3, 2} },
		},
		{
			2, 0,
			data.Table{
				&data.SliceInt{1, 2, 3, 4},
				&data.SliceFloat32{.1, .2, .3, .4},
			},
			data.Table{
				&data.SliceInt{3, 2, 1, 4},
				&data.SliceFloat32{.3, .2, .1, .4},
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
	cases := []struct{
		i      int
		table  data.Table
		result data.Table
	}{
		{
			0,
			data.Table{},
			data.Table{},
		},
		{
			0,
			data.Table{
				&data.SliceInt{1, 2, 3, 4},
				&data.SliceFloat32{1, 2, 3, 4},
			},
			data.Table{
				&data.SliceInt{4, 2, 3},
				&data.SliceFloat32{4, 2, 3},
			},
		},
		{
			1,
			data.Table{
				&data.SliceInt{1, 2, 3, 4},
				&data.SliceFloat32{1, 2, 3, 4},
			},
			data.Table{
				&data.SliceInt{1, 4, 3},
				&data.SliceFloat32{1, 4, 3},
			},
		},
		{
			3,
			data.Table{
				&data.SliceInt{1, 2, 3, 4},
				&data.SliceFloat32{1, 2, 3, 4},
			},
			data.Table{
				&data.SliceInt{1, 2, 3},
				&data.SliceFloat32{1, 2, 3},
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
