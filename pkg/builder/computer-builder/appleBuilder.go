package computer_builder

import (
	computer2 "github.com/VitalyDorozhkin/go-patterns/pkg/builder/computer"
)

type appleBuilder struct {
	computer *computer2.Computer
}

func (a *appleBuilder) SetType(model string) {
	a.computer.Type = model
}

func (a *appleBuilder) SetMotherBoard(model string) {
	a.computer.MotherBoard = "Apple " + model + " MotherBoard"
}

func (a *appleBuilder) SetCPU(model string) {
	a.computer.CPU = model
}

func (a *appleBuilder) SetHardDrive(model string) {
	a.computer.HardDrive = "Apple HardDrive on " + model
}

func (a *appleBuilder) SetRam(model string) {
	a.computer.Ram = "Apple Ram on " + model
}

func (a *appleBuilder) SetGPU(model string) {
	a.computer.GPU = model
}

func (a *appleBuilder) SetOS(model string) {
	a.computer.OS = "MacOS Catalina " + model
}

func (a *appleBuilder) Build() *computer2.Computer {
	return a.computer
}

func NewAppleBuilder() ComputerBuilder{
	return &appleBuilder{&computer2.Computer{Builder: "Apple"}}
}
