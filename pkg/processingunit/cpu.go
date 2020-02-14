package processingunit

import (
	"github.com/VitalyDorozhkin/go-patterns/pkg/datastorage"
)

type CPU struct {
	pointer int
	cache   byte
}

func (c *CPU) Freeze() {
	c.pointer = 0
}

func (c *CPU) SetCache(memory datastorage.Memory) {
	c.cache = memory.Read(c.pointer)
}

func (c *CPU) LoadCache(memory datastorage.Memory) {
	memory.Write(c.pointer, c.cache)
}

func (c *CPU) Jump(position int) {
	c.pointer = position
}

func (c *CPU) Execute() int {
	return c.pointer
}
