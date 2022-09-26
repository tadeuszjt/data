package dataTest

import (
	"testing"
	. "github.com/tadeuszjt/data"
)

func BenchmarkTableLen(b *testing.B) {
	table := Table{
		&RowT[int]{1, 2, 3, 4},
	}
	
	for i := 0; i < b.N; i++ {
		table.Len()
	}
}

func BenchmarkTableSwap(b *testing.B) {
	table := Table{
		&RowT[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
	}
	
	for i := 0; i < b.N; i++ {
		table.Swap(i % 16, (i*7) % 16)
	}
}

func BenchmarkTableDelete(b *testing.B) {
	row := RowT[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	var row2 RowT[int] = row[:]
	table := Table{ &row2 }
	
	for i := 0; i < b.N; i++ {
		row2 = row[:]
		table.Delete(i % 16)
	}
}

