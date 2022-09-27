package dataTest

import (
	. "github.com/tadeuszjt/data"
	"testing"
)

func keyMapIdentical(a, b KeyMap) bool {
	if len(a.KeyToIndex) != len(b.KeyToIndex) {
		return false
	}
	for i := range a.KeyToIndex {
		if a.KeyToIndex[i] != b.KeyToIndex[i] {
			return false
		}
	}

	aRow, ok := a.Row.(*RowT[int])
	if !ok {
		panic("!ok")
	}

	bRow, ok := b.Row.(*RowT[int])
	if !ok {
		panic("!ok")
	}

	return rowIntIdentical(*aRow, *bRow)
}

func TestKeyMapLen(t *testing.T) {
	cases := []struct {
		km     KeyMap
		result int
	}{
		{
			KeyMap{
				&RowT[int]{},
				[]int{},
			},
			0,
		},
		{
			KeyMap{
				&RowT[int]{1, 2, 3},
				[]int{},
			},
			3,
		},
		{
			KeyMap{
				&RowT[int]{1, 2, 3, 4},
				[]int{1, 2, 3},
			},
			4,
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.km.Len()

		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestKeyMapAppend(t *testing.T) {
	cases := []struct {
		keyMap KeyMap
		items  []any
		result KeyMap
		ret    Key
	}{
		{
			KeyMap{
				&RowT[int]{},
				[]int{},
			},
			[]any{1},
			KeyMap{
				&RowT[int]{1},
				[]int{0},
			},
			Key(0),
		},
		{
			KeyMap{
				&RowT[int]{},
				[]int{-1},
			},
			[]any{2},
			KeyMap{
				&RowT[int]{2},
				[]int{0},
			},
			Key(0),
		},
		{
			KeyMap{
				&RowT[int]{1, 2, 3},
				[]int{1, 0},
			},
			[]any{5},
			KeyMap{
				&RowT[int]{1, 2, 3, 5},
				[]int{1, 0, 3},
			},
			Key(2),
		},
		{
			KeyMap{
				&RowT[int]{1, 2, 3},
				[]int{1, -1},
			},
			[]any{5},
			KeyMap{
				&RowT[int]{1, 2, 3, 5},
				[]int{1, 3},
			},
			Key(1),
		},
	}

	for _, c := range cases {
		expected := c.result
		ret := c.keyMap.Append(c.items...)
		actual := c.keyMap

		if !keyMapIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
		if ret != c.ret {
			t.Errorf("ret: expected: %v, actual: %v", c.ret, ret)
		}
	}
}
func TestKeyMapDelete(t *testing.T) {
	cases := []struct {
        keyMap KeyMap
        key    Key
		result KeyMap
	}{
		{
			KeyMap{
				&RowT[int]{1},
				[]int{0},
			},
            Key(0),
			KeyMap{
				&RowT[int]{},
				[]int{},
			},
		},
		{
			KeyMap{
				&RowT[int]{1, 2},
				[]int{0, 1},
			},
            Key(0),
			KeyMap{
				&RowT[int]{2},
				[]int{-1, 0},
			},
		},
		{
			KeyMap{
				&RowT[int]{1, 2},
				[]int{0, 1},
			},
            Key(1),
			KeyMap{
				&RowT[int]{1},
				[]int{0},
			},
		},
	}

	for _, c := range cases {
		expected := c.result
        c.keyMap.Delete(c.key)
		actual := c.keyMap

		if !keyMapIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

