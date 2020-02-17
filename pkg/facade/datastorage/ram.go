package datastorage

import "fmt"

type Ram interface {
	Read(position int) byte
	Write(position int, data byte)
	CleanAll()
	Storage() Storage
}

type ram struct {
	storage
}

func (r *ram) Storage() Storage {
	return &r.storage
}

func (r *ram) Read(position int) byte {
	if position >= r.Size() {
		fmt.Printf("Ram don't has %d pointer\n", position)
		return 0
	}
	fmt.Printf("Ram read from %d\n", position)
	return r.disc[position]
}

func (r *ram) Write(position int, data byte) {
	if position >= r.Size() {
		fmt.Printf("Ram don't has %d pointer\n", position)
		return
	}
	fmt.Printf("Ram write %b to %d\n", data, position)
	r.disc[position] = data
}

func (r *ram) CleanAll() {
	for i := range r.head {
		r.head[i] = false
	}
}

func NewRam() Ram {
	return &ram{}
}
