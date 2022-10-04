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
	if a.Len() != b.Len() {
		return false
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
		{MakeKeyMap(&RowT[int]{}), 0},
		{MakeKeyMap(&RowT[int]{1, 2, 3}), 3},
		{MakeKeyMap(&RowT[int]{1, 2, 3, 4}), 4},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.km.Len()

		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
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
	CheckEqual(t, 0, keyMap.GetIndex(key0))
	CheckEqual(t, 1, keyMap.GetIndex(key1))
	CheckEqual(t, 0, row[0])
	CheckEqual(t, 1, row[1])

	// Row       : [0, 1]
	// KeyToIndex: [0, 1]
	// unusedKeys: []

	key2 = keyMap.Append(2)
	CheckEqual(t, 3, keyMap.Len())
	CheckEqual(t, 0, keyMap.GetIndex(key0))
	CheckEqual(t, 1, keyMap.GetIndex(key1))
	CheckEqual(t, 2, keyMap.GetIndex(key2))
	CheckEqual(t, 0, row[0])
	CheckEqual(t, 1, row[1])
	CheckEqual(t, 2, row[2])

	// Row       : [0, 1, 2]
	// KeyToIndex: [0, 1, 2]
	// unusedKeys: []

	keyMap.Delete(key1)
	CheckEqual(t, 2, keyMap.Len())
	CheckEqual(t, 0, keyMap.GetIndex(key0))
	CheckEqual(t, 1, keyMap.GetIndex(key2))
	CheckEqual(t, 0, row[0])
	CheckEqual(t, 2, row[1])

	// Row       : [0, 2]
	// KeyToIndex: [0, -1, 2]
	// unusedKeys: [1]

	keyMap.Delete(key2)
	CheckEqual(t, 1, keyMap.Len())
	CheckEqual(t, 0, keyMap.GetIndex(key0))
	CheckEqual(t, 0, row[0])

	// Row       : [0]
	// KeyToIndex: [0, -1]
	// unusedKeys: [1]

	key1 = keyMap.Append(1)
	CheckEqual(t, 2, keyMap.Len())
	CheckEqual(t, 0, keyMap.GetIndex(key0))
	CheckEqual(t, 1, keyMap.GetIndex(key1))
	CheckEqual(t, 0, row[0])
	CheckEqual(t, 1, row[1])

	// Row       : [0, 1]
	// KeyToIndex: [0, 1]
	// unusedKeys: []
}
