package dataTest

import (
	. "github.com/tadeuszjt/data"
	"testing"
)

func CheckEqual[T ~int | ~bool](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

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
		item   any
		result KeyMap
		ret    Key
	}{
		{
			MakeKeyMap(&RowT[int]{}),
			1,
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
			5,
			KeyMap{
				Row:        &RowT[int]{1, 2, 3, 5},
				KeyToIndex: []int{1, 0, 3},
			},
			Key(2),
		},
	}

	for _, c := range cases {
		expected := c.result
		ret := c.keyMap.Append(c.item)
		actual := c.keyMap

		if !keyMapIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
		if ret != c.ret {
			t.Errorf("ret: expected: %v, actual: %v", c.ret, ret)
		}
	}
}

func TestKeyMapDeleteAppend(t *testing.T) {
	row := RowT[int]{}
	keyMap := MakeKeyMap(&row)

	key0 := keyMap.Append(0)
	key1 := keyMap.Append(1)
	key2 := keyMap.Append(2)

	CheckEqual(t, 0, key0)
	CheckEqual(t, 1, key1)
	CheckEqual(t, 2, key2)
	CheckEqual(t, 3, keyMap.Len())
	CheckEqual(t, 0, row[0])
	CheckEqual(t, 1, row[1])
	CheckEqual(t, 2, row[2])

	// Row       : [0, 1, 2]
	// KeyToIndex: [0, 1, 2]
	// unusedKeys: []

	keyMap.Delete(key2)
	CheckEqual(t, 2, keyMap.Len())
	CheckEqual(t, 0, keyMap.KeyToIndex[key0])
	CheckEqual(t, 1, keyMap.KeyToIndex[key1])
	CheckEqual(t, 0, row[0])
	CheckEqual(t, 1, row[1])

	// Row       : [0, 1]
	// KeyToIndex: [0, 1]
	// unusedKeys: []

	key2 = keyMap.Append(2)
	CheckEqual(t, 3, keyMap.Len())
	CheckEqual(t, 0, keyMap.KeyToIndex[key0])
	CheckEqual(t, 1, keyMap.KeyToIndex[key1])
	CheckEqual(t, 2, keyMap.KeyToIndex[key2])
	CheckEqual(t, 0, row[0])
	CheckEqual(t, 1, row[1])
	CheckEqual(t, 2, row[2])

	// Row       : [0, 1, 2]
	// KeyToIndex: [0, 1, 2]
	// unusedKeys: []

	keyMap.Delete(key1)
	CheckEqual(t, 2, keyMap.Len())
	CheckEqual(t, 0, keyMap.KeyToIndex[key0])
	CheckEqual(t, -1, keyMap.KeyToIndex[key1])
	CheckEqual(t, 1, keyMap.KeyToIndex[key2])
	CheckEqual(t, 0, row[0])
	CheckEqual(t, 2, row[1])

	// Row       : [0, 2]
	// KeyToIndex: [0, -1, 2]
	// unusedKeys: [1]

	keyMap.Delete(key2)
	CheckEqual(t, 1, keyMap.Len())
	CheckEqual(t, 0, keyMap.KeyToIndex[key0])
	CheckEqual(t, -1, keyMap.KeyToIndex[key1])
	CheckEqual(t, 0, row[0])
	CheckEqual(t, 2, len(keyMap.KeyToIndex))

	// Row       : [0]
	// KeyToIndex: [0, -1]
	// unusedKeys: [1]

	key1 = keyMap.Append(1)
	CheckEqual(t, 2, keyMap.Len())
	CheckEqual(t, 0, keyMap.KeyToIndex[key0])
	CheckEqual(t, 1, keyMap.KeyToIndex[key1])
	CheckEqual(t, 0, row[0])
	CheckEqual(t, 1, row[1])
	CheckEqual(t, 2, len(keyMap.KeyToIndex))

	// Row       : [0, 1]
	// KeyToIndex: [0, 1]
	// unusedKeys: []
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
