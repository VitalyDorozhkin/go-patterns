package cpu

import (
	"fmt"
	"github.com/VitalyDorozhkin/go-patterns/pkg/facade/memory"
	)
type CPU struct {
	pointer int
	cache   byte
}

func (c *CPU) Freeze() {
	fmt.Println("CPU freeze")
	c.pointer = 0
}

func (c *CPU) SetCache(memory memory.Memory) {
	fmt.Println("CPU write cache")
	c.cache = memory.Read(c.pointer)
}

func (c *CPU) LoadCache(memory memory.Memory) {
	fmt.Println("CPU load cache to memory")
	memory.Write(c.pointer, c.cache)
}

func (c *CPU) Jump(position int) {
	fmt.Printf("CPU jump to %d\n", position)
	c.pointer = position
}

func (c *CPU) Execute() int {
	fmt.Println("CPU execute")
	return c.pointer
}
