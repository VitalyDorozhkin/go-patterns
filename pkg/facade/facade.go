package facade

import (
	"fmt"
	"github.com/VitalyDorozhkin/go-patterns/pkg/facade/datastorage"
	"github.com/VitalyDorozhkin/go-patterns/pkg/facade/processingunit"
)

type ComputerFacade struct {
	processor *processingunit.CPU
	ram       *datastorage.Ram
	hd        *datastorage.HardDrive
}

func (c *ComputerFacade) Start() int {
	c.processor.Freeze()
	c.ram.Write(0, c.hd.Read(0))
	c.processor.Jump(8)
	return c.processor.Execute()
}

func (c *ComputerFacade) Swap(ramPos int, drivePos int) {
	c.processor.Jump(ramPos)
	c.processor.SetCache(c.ram)
	c.processor.Jump(drivePos)
	c.processor.LoadCache(c.hd)
}

func (c *ComputerFacade) TurnOff() {
	for i := 0; i < c.ram.Size(); i++ {
		c.Swap(i, i)
	}
	c.processor.Freeze()
	fmt.Println("Bye!")
}

func (c *ComputerFacade) Write(data byte) (position int) {
	position = c.hd.WriteFirst(data)
	if position < c.hd.Size() {
		fmt.Printf("Data loaded on HardDrive[%d]\n", position)
		return
	}
	fmt.Println("Hard drive is full!")
	return
}

func NewComputerFacade() *ComputerFacade {

	c := &ComputerFacade{
		processor: &processingunit.CPU{},
		ram:       &datastorage.Ram{},
		hd:        &datastorage.HardDrive{},
	}
	c.ram.SetSize(4)
	c.hd.SetSize(16)
	return c
}
