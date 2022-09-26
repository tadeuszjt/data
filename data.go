package data

type Row interface {
	Len() int
	Swap(i, j int)
	Delete(i int)
	Append(t any)
}

/*
 * A Row represents a contiguous dynamically-allocated slice of data which can be controlled using
 * the interface functions.
 */
type RowT[T any] []T

func (s *RowT[T]) Len() int {
	return len(*s)
}

func (s *RowT[T]) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *RowT[T]) Delete(i int) {
	(*s)[i] = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
}

func (s *RowT[T]) Append(t any) {
	i, ok := t.(T)
	if !ok {
		panic("wrong type")
	}

	*s = append(*s, i)
}


/*
 * A table is a collection of rows which are all controlled simultaneously.
 */
type Table []Row

func (t Table) Len() int {
	return t[0].Len()
}

func (t Table) Swap(i, j int) {
	for k := range t {
		t[k].Swap(i, j)
	}
}

func (t Table) Delete(i int) {
	for k := range t {
		t[k].Delete(i)
	}
}

func (t Table) Append(items ...any) {
	for i, item := range items {
		t[i].Append(item)
	}
}

func (t Table) Filter(f func(int) bool) {
	for i := 0; i < t.Len(); i++ {
		if !f(i) {
			t.Delete(i)
			i--
		}
	}
}


/*
 * A keymap is a device which returns integer keys when allocating elements in a table.
 * Provided the table state has not been modified by other functions, the keys can be used to
 * always retrieve the same associated elements even when the table columns have been reordered
 * due to calls to Append()/Delete().
 */
type KeyMap struct {
    Table *Table
    KeyToIndex []int
}

func (k *KeyMap) Len() int {
    return k.Table.Len()
}
