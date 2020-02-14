package main

import (
	"fmt"
	"github.com/VitalyDorozhkin/go-patterns/pkg/facade"
)

func main() {
	computer := facade.NewComputerFacade()
	computer.Start()
	position := computer.Write(11)
	fmt.Printf("position: %d\n", position)
	computer.TurnOff()
}
