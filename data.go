package data

type Slice interface{
	Len() int
	Swap(i, j int)
	Delete(i int)
}

type Table []Slice

func (t *Table) Len() int {
	return (*t)[0].Len()
}

func (t *Table) Swap(i, j int) {
	for k := range *t {
		(*t)[k].Swap(i, j)
	}
}

func (t *Table) Delete(i int) {
	for k := range *t {
		(*t)[k].Delete(i)
	}
}
