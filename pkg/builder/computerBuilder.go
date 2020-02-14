package builder

type ComputerBuilder interface {
	InitComputer()
	SetType(model string)
	SetMotherBoard(model string)
	SetCPU(model string)
	SetHardDrive(model string)
	SetRam(model string)
	SetGPU(model string)
	SetOS(model string)
	Build() *Computer
}
