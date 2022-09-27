package dataTest

import (
	"github.com/tadeuszjt/data"
	"testing"
)

func BenchmarkKeyMap(b *testing.B) {
    keyMap := data.KeyMap { Row: &data.RowT[int]{} }
    keys := []data.Key{}

    numKeys := 20000

    for i := 0; i < numKeys; i++ {
        keys = append(keys, keyMap.Append(i))
    }

	for i := 0; i < b.N; i++ {
        x := (i * 7) % numKeys
        keyMap.Delete(keys[x])
        keys[x] = keyMap.Append(i)
	}
}
