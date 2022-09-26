package dataTest

import (
	"testing"
	. "github.com/tadeuszjt/data"
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
	
	return tableIdentical(*a.Table, *b.Table)
}


func TestKeyMapLen(t *testing.T) {
	cases := []struct{
		km     KeyMap
		result int
	}{
		{
            KeyMap {
                &Table{ &RowT[int]{} },
                []int{},
            },
            0,
		},
		{
            KeyMap {
                &Table{ &RowT[int]{ 1, 2, 3 } },
                []int{},
            },
            3,
		},
		{
            KeyMap {
                &Table{
                    &RowT[int]{ 1, 2, 3, 4},
                    &RowT[float32]{3, 4, 5, 6}},
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

//func TestKeyMapSwap(t *testing.T) {
//	cases := []struct{
//        i, j   int
//		km     KeyMap
//		result KeyMap
//	}{
//		{
//            0, 0,
//            KeyMap {
//                &Table{ &RowT[int]{} },
//                []int{},
//            },
//            KeyMap {
//                &Table{ &RowT[int]{} },
//                []int{},
//            },
//		},
//	}
//	
//	for _, c := range cases {
//        c.km.Swap(c.i, c.j)
//
//		if !keyMapIdentical(c.km, c.result) {
//			t.Errorf("expected: %v, actual: %v", c.result, c.km)
//		}
//	}
//}
