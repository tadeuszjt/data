package data

const (
	KeyInvalid = Key(-1)
)

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
	keyToIndex RowT[int]
	unusedKeys RowT[Key]
}

func MakeKeyMap(row Row) KeyMap {
	return KeyMap{Row: row}
}

func (k *KeyMap) GetIndex(key Key) int {
	if key < 0 || key > Key(len(k.keyToIndex)-1) || k.keyToIndex[key] < 0 || k.keyToIndex[key] > k.Row.Len() {
		panic("invalid key")
	}
	return k.keyToIndex[key]
}

func (k *KeyMap) Len() int {
	return k.Row.Len()
}

func (k *KeyMap) Append(items ...any) Key {
	if k.unusedKeys.Len() > 0 { // use empty slot
		key := k.unusedKeys[k.unusedKeys.Len()-1]
		k.unusedKeys.Delete(k.unusedKeys.Len() - 1)

		k.keyToIndex[key] = k.Row.Len()
		k.Row.Append(items...)
		return key
	}

	// allocate new slot
	index := k.Row.Len()
	key := k.keyToIndex.Len()
	k.Row.Append(items...)
	k.keyToIndex.Append(index)
	return Key(key)
}

func (k *KeyMap) Delete(key Key) {
	index := k.GetIndex(key)

	end := k.Row.Len() - 1
	if index != end { // swap row elements
		for i := range k.keyToIndex {
			if k.keyToIndex[i] == end {
				k.keyToIndex[i] = index
				break
			}
		}
	}

	k.Row.Delete(index)

	if key == Key(k.keyToIndex.Len()-1) { // key points to end element
		k.keyToIndex.Delete(k.keyToIndex.Len() - 1)
	} else {
		k.keyToIndex[key] = -1
		k.unusedKeys.Append(key)
	}
}
