package index07

type SubGramOffset struct {
	subGram string
	offset  int
}

func (s *SubGramOffset) SubGram() string {
	return s.subGram
}

func (s *SubGramOffset) SetSubGram(subGram string) {
	s.subGram = subGram
}

func (s *SubGramOffset) Offset() int {
	return s.offset
}

func (s *SubGramOffset) SetOffset(offset int) {
	s.offset = offset
}

func NewSubGramOffset(subGram string, offset int) SubGramOffset {
	return SubGramOffset{
		subGram: subGram,
		offset:  offset,
	}
}
