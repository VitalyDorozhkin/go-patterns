package datastorage

import "fmt"

type HardDrive interface {
	Read(position int) byte
	Write(position int, data byte)
	WriteFirst(data byte) (i int)
	Storage() Storage
}

type hardDrive struct {
	storage
}

func (h *hardDrive) Storage() Storage {
	return &h.storage
}

func (h *hardDrive) Read(position int) byte {
	if position >= h.Size() {
		fmt.Printf("HardDrive don't has %d pointer\n", position)
		return 0
	}
	fmt.Printf("HardDrive read from %d\n", position)
	return h.disc[position]
}

func (h *hardDrive) Write(position int, data byte) {
	if position >= h.Size() {
		fmt.Printf("HardDrive don't has %d pointer\n", position)
		return
	}
	fmt.Printf("HardDrive write %b to %d\n", data, position)
	h.disc[position] = data
}

func (h *hardDrive) WriteFirst(data byte) (i int) {
	for i = 0; i < h.Size(); i++ {
		if !h.head[i] {
			h.disc[i] = data
			h.head[i] = true
			break
		}
	}
	return i
}

func NewHardDrive() HardDrive {
	return &hardDrive{}
}
