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
				Row:        &RowT[int]{},
				KeyToIndex: []int{},
			},
			0,
		},
		{
			KeyMap{
				Row:        &RowT[int]{1, 2, 3},
				KeyToIndex: []int{},
			},
			3,
		},
		{
			KeyMap{
				Row:        &RowT[int]{1, 2, 3, 4},
				KeyToIndex: []int{1, 2, 3},
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
			MakeKeyMap(&RowT[int]{}),
			[]any{1},
			KeyMap{
				Row:        &RowT[int]{1},
				KeyToIndex: []int{0},
			},
			Key(0),
		},
		{
			KeyMap{
				Row:        &RowT[int]{1, 2, 3},
				KeyToIndex: []int{1, 0},
			},
			[]any{5},
			KeyMap{
				Row:        &RowT[int]{1, 2, 3, 5},
				KeyToIndex: []int{1, 0, 3},
			},
			Key(2),
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
				Row:        &RowT[int]{1},
				KeyToIndex: []int{0},
			},
			Key(0),
			KeyMap{
				Row:        &RowT[int]{},
				KeyToIndex: []int{},
			},
		},
		{
			KeyMap{
				Row:        &RowT[int]{1, 2},
				KeyToIndex: []int{0, 1},
			},
			Key(0),
			KeyMap{
				Row:        &RowT[int]{2},
				KeyToIndex: []int{-1, 0},
			},
		},
		{
			KeyMap{
				Row:        &RowT[int]{1, 2},
				KeyToIndex: []int{0, 1},
			},
			Key(1),
			KeyMap{
				Row:        &RowT[int]{1},
				KeyToIndex: []int{0},
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
