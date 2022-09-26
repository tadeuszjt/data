package data


type RowT[T any] []T


func (s *RowT[T]) Len() int {
    return len(*s)
}


func (s *RowT[T]) Swap(i, j int) {
    (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *RowT[T]) Delete(i int) {
    (*s)[i] = (*s)[len(*s) - 1]
    *s = (*s)[:len(*s) - 1]
}

func (s *RowT[T]) Append(t interface{}) {
	i, ok := t.(T)
	if !ok {
		panic("wrong type")
	}
	
	*s = append(*s, i)
}


type Row interface{
	Len() int
	Swap(i, j int)
	Delete(i int)
	Append(t interface{})
}

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

func (t Table) Append(items ...interface{}) {
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
