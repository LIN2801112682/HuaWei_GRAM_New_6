package matchQuery

import "index07"

type SortKey struct {
	offset             int
	sizeOfInvertedList int
	gram               string
	invertedIndex      index07.Inverted_index
}

func (s *SortKey) Offset() int {
	return s.offset
}

func (s *SortKey) SetOffset(offset int) {
	s.offset = offset
}

func (s *SortKey) SizeOfInvertedList() int {
	return s.sizeOfInvertedList
}

func (s *SortKey) SetSizeOfInvertedList(sizeOfInvertedList int) {
	s.sizeOfInvertedList = sizeOfInvertedList
}

func (s *SortKey) Gram() string {
	return s.gram
}

func (s *SortKey) SetGram(gram string) {
	s.gram = gram
}

func (s *SortKey) InvertedIndex() index07.Inverted_index {
	return s.invertedIndex
}

func (s *SortKey) SetInvertedIndex(invertedIndex index07.Inverted_index) {
	s.invertedIndex = invertedIndex
}

func NewSortKey(offset int, pos int, gram string, invertIndex index07.Inverted_index) SortKey {
	return SortKey{
		offset:             offset,
		sizeOfInvertedList: pos,
		gram:               gram,
		invertedIndex:      invertIndex,
	}
}
