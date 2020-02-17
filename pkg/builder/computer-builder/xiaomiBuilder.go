package computer_builder

import computer2 "github.com/VitalyDorozhkin/go-patterns/pkg/builder/computer"

type xiaomiBuilder struct {
	computer *computer2.Computer
}

func (x *xiaomiBuilder) SetType(model string) {
	x.computer.Type = model
}

func (x *xiaomiBuilder) SetMotherBoard(model string) {
	x.computer.MotherBoard = "Xiaomi " + model + " MotherBoard"
}

func (x *xiaomiBuilder) SetCPU(model string) {
	x.computer.CPU = model
}

func (x *xiaomiBuilder) SetHardDrive(model string) {
	x.computer.HardDrive = "Samsung HardDrive on " + model
}

func (x *xiaomiBuilder) SetRam(model string) {
	x.computer.Ram = "Samsung Ram on " + model
}

func (x *xiaomiBuilder) SetGPU(model string) {
	x.computer.GPU = model
}

func (x *xiaomiBuilder) SetOS(model string) {
	x.computer.OS = "Windows 10 " + model
}

func (x *xiaomiBuilder) Build() *computer2.Computer {
	return x.computer
}

func NewXiaomiBuilder() ComputerBuilder{
	return &xiaomiBuilder{&computer2.Computer{Builder: "Xiaomi"}}
}
