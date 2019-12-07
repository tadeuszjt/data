package data

type SliceFloat32 []float32

func (s *SliceFloat32) Len() int {
	return len(*s)
}

func (s *SliceFloat32) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *SliceFloat32) Delete(i int) {
	end := s.Len() - 1
	if i < end {
		s.Swap(i, end)
	}
	
	*s = (*s)[:end]
}

