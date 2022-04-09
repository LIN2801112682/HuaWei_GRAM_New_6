package matchQuery

type SortKey struct {
	sizeOfInvertedList int
	gram               string
}

func (s *SortKey) Pos() int {
	return s.sizeOfInvertedList
}

func (s *SortKey) SetPos(pos int) {
	s.sizeOfInvertedList = pos
}

func (s *SortKey) Gram() string {
	return s.gram
}

func (s *SortKey) SetGram(gram string) {
	s.gram = gram
}

func NewSortKey(pos int, gram string) SortKey {
	return SortKey{
		sizeOfInvertedList: pos,
		gram:               gram,
	}
}
