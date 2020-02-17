package builder

type XiaomiBuilder struct {
	computer *Computer
}

func (x *XiaomiBuilder) InitComputer() {
	x.computer = new(Computer)
	x.computer.Builder = "Xiaomi"
}

func (x *XiaomiBuilder) SetType(model string) {
	x.computer.Type = model
}

func (x *XiaomiBuilder) SetMotherBoard(model string) {
	x.computer.MotherBoard = "Xiaomi " + model + " MotherBoard"
}

func (x *XiaomiBuilder) SetCPU(model string) {
	x.computer.CPU = model
}

func (x *XiaomiBuilder) SetHardDrive(model string) {
	x.computer.HardDrive = "Samsung HardDrive on " + model
}

func (x *XiaomiBuilder) SetRam(model string) {
	x.computer.Ram = "Samsung Ram on " + model
}

func (x *XiaomiBuilder) SetGPU(model string) {
	x.computer.GPU = model
}

func (x *XiaomiBuilder) SetOS(model string) {
	x.computer.OS = "Windows 10 " + model
}

func (x *XiaomiBuilder) Build() *Computer {
	return x.computer
}
