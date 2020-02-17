package processingunit

import (
	"github.com/VitalyDorozhkin/go-patterns/pkg/facade/datastorage"
)

type CPU interface {
	Freeze()
	SetCache(memory datastorage.MemoryEditor)
	LoadCache(memory datastorage.MemoryEditor)
	Jump(position int)
	Execute() int
}

type cpu struct {
	pointer int
	cache   byte
}

func (c *cpu) Freeze() {
	c.pointer = 0
}

func (c *cpu) SetCache(memory datastorage.MemoryEditor) {
	c.cache = memory.Read(c.pointer)
}

func (c *cpu) LoadCache(memory datastorage.MemoryEditor) {
	memory.Write(c.pointer, c.cache)
}

func (c *cpu) Jump(position int) {
	c.pointer = position
}

func (c *cpu) Execute() int {
	return c.pointer
}

func NewCPU() CPU {
	return &cpu{}
}
