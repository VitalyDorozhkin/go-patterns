package computer_builder

import (
	"github.com/VitalyDorozhkin/go-patterns/pkg/builder/computer"
)

type ComputerBuilder interface {
	SetType(model string)
	SetMotherBoard(model string)
	SetCPU(model string)
	SetHardDrive(model string)
	SetRam(model string)
	SetGPU(model string)
	SetOS(model string)
	Build() *computer.Computer
}
