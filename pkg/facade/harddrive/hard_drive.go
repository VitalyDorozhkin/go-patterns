package harddrive

import "fmt"

type HardDrive struct {
	disc [16]byte
	head [16]bool
}

func (h *HardDrive) Free(position int) {
	if position < h.GetSize() {
		h.head[position] = false
	}
}

func (h *HardDrive) CleanAll() {
	for i := range h.head {
		h.head[i] = false
	}
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

func (h *HardDrive) GetSize() int {
	return len(h.disc)
}

