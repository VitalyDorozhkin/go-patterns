package director

import (
	"github.com/VitalyDorozhkin/go-patterns/pkg/builder/computer"
	"github.com/VitalyDorozhkin/go-patterns/pkg/builder/computer-builder"
)

type Director interface {
	ConstructPro() *computer.Computer
	ConstructAir() *computer.Computer
}

type director struct {
	builder computer_builder.ComputerBuilder
}

func (d *director) ConstructPro() *computer.Computer {
	d.builder.SetType("Pro")
	d.builder.SetMotherBoard("Pro")
	d.builder.SetCPU("Intel Xeon W12")
	d.builder.SetHardDrive("1 TB")
	d.builder.SetRam("96 GB")
	d.builder.SetGPU("AMD Radeon Pro 580X")
	d.builder.SetOS("Pro")
	return d.builder.Build()
}

func (d *director) ConstructAir() *computer.Computer {
	d.builder.SetType("Air")
	d.builder.SetMotherBoard("Air")
	d.builder.SetCPU("Intel Core i5")
	d.builder.SetGPU("Nvidia Geforce")
	d.builder.SetHardDrive("500 GB")
	d.builder.SetRam("16 GB")
	d.builder.SetOS("Lite")
	return d.builder.Build()
}

func NewService(builder computer_builder.ComputerBuilder) Director{
	return &director{builder}
}
