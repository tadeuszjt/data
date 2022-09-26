package data

type Row interface {
	Len() int
	Swap(i, j int)
	Delete(i int)
	Append(items ...any)
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

func (s *RowT[T]) Append(items ...any) {
	if len(items) != 1 {
		panic("Num items != 1")
	}

	item, ok := items[0].(T)
	if !ok {
		panic("wrong type")
	}

	*s = append(*s, item)
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
type Key int

type KeyMap struct {
	Row        Row
	KeyToIndex []int
}

func (k *KeyMap) Len() int {
	return k.Row.Len()
}

func (k *KeyMap) Append(items ...any) Key {
	for i := range k.KeyToIndex {
		if k.KeyToIndex[i] < 0 { // use empty slot
			k.KeyToIndex[i] = k.Row.Len()
			k.Row.Append(items)
			return Key(i)
		}
	}

	// allocate new slot
	index := k.Row.Len()
	key := len(k.KeyToIndex)
	k.Row.Append(items)
	k.KeyToIndex = append(k.KeyToIndex, index)
	return Key(key)
}

func (k *KeyMap) Delete(key Key) {
	index := k.KeyToIndex[key]

	end := k.Row.Len() - 1
	if index != end { // swap row elements
		for i := range k.KeyToIndex {
			if k.KeyToIndex[i] == end {
				k.KeyToIndex[i] = index
				break
			}
		}
	}

	k.Row.Delete(index)

	if key == Key(len(k.KeyToIndex)-1) { // key points to end element
		k.KeyToIndex = k.KeyToIndex[:len(k.KeyToIndex)-1]
	} else {
		k.KeyToIndex[key] = -1
	}
}
