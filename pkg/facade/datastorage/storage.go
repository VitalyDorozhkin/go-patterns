package datastorage

type storage struct {
	disc []byte
	head []bool
}

type Storage interface {
	Free(position int)
	Size() int
	CleanAll()
	SetSize(size int)
}

func (s *storage) Free(position int) {
	if position < s.Size() {
		s.head[position] = false
	}
}

func (s *storage) CleanAll() {
	for i := range s.head {
		s.head[i] = false
	}
}

func (s *storage) Size() int {
	return len(s.disc)
}

func (s *storage) SetSize(size int) {
	s.disc = make([]byte, size, size)
	s.head = make([]bool, size, size)
}

