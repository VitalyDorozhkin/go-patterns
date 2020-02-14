package datastorage

type Disc struct {
	disc []byte
	head []bool
}

func (d *Disc) Free(position int) {
	if position < d.GetSize() {
		d.head[position] = false
	}
}

func (d *Disc) CleanAll() {
	for i := range d.head {
		d.head[i] = false
	}
}

func (d *Disc) GetSize() int {
	return len(d.disc)
}

func NewDisc(size int) *Disc {
	return &Disc{make([]byte, size, size), make([]bool, size, size)}
}
