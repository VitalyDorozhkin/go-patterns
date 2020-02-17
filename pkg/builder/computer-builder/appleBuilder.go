package builder

type AppleBuilder struct {
	computer *Computer
}

func (a *AppleBuilder) InitComputer() {
	a.computer = new(Computer)
	a.computer.Builder = "Apple"
}

func (a *AppleBuilder) SetType(model string) {
	a.computer.Type = model
}

func (a *AppleBuilder) SetMotherBoard(model string) {
	a.computer.MotherBoard = "Apple " + model + " MotherBoard"
}

func (a *AppleBuilder) SetCPU(model string) {
	a.computer.CPU = model
}

func (a *AppleBuilder) SetHardDrive(model string) {
	a.computer.HardDrive = "Apple HardDrive on " + model
}

func (a *AppleBuilder) SetRam(model string) {
	a.computer.Ram = "Apple Ram on " + model
}

func (a *AppleBuilder) SetGPU(model string) {
	a.computer.GPU = model
}

func (a *AppleBuilder) SetOS(model string) {
	a.computer.OS = "MacOS Catalina " + model
}

func (a *AppleBuilder) Build() *Computer {
	return a.computer
}
