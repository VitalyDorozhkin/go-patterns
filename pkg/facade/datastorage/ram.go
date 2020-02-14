package datastorage

import "fmt"

type Ram struct {
	storage
}

func (r *Ram) Read(position int) byte {
	if position >= r.Size() {
		fmt.Printf("Ram don't has %d pointer\n", position)
		return 0
	}
	fmt.Printf("Ram read from %d\n", position)
	return r.disc[position]
}

func (r *Ram) Write(position int, data byte) {
	if position >= r.Size() {
		fmt.Printf("Ram don't has %d pointer\n", position)
		return
	}
	fmt.Printf("Ram write %b to %d\n", data, position)
	r.disc[position] = data
}

func (r *Ram) CleanAll() {
	for i := range r.head {
		r.head[i] = false
	}
}
