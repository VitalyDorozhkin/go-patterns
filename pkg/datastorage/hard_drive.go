package datastorage

import "fmt"

type HardDrive struct {
	Disc
}

func (h *HardDrive) Read(position int) byte {
	if position >= h.GetSize() {
		fmt.Printf("HardDrive don't has %d pointer\n", position)
		return 0
	}
	fmt.Printf("HardDrive read from %d\n", position)
	return h.disc[position]
}

func (h *HardDrive) Write(position int, data byte) {
	if position >= h.GetSize() {
		fmt.Printf("HardDrive don't has %d pointer\n", position)
		return
	}
	fmt.Printf("HardDrive write %b to %d\n", data, position)
	h.disc[position] = data
}

func (h *HardDrive) WriteFirst(data byte) (i int) {
	for i = 0; i < h.GetSize(); i++ {
		if !h.head[i] {
			h.disc[i] = data
			h.head[i] = true
			break
		}
	}
	return i
}
