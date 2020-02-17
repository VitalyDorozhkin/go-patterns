package main

import (
	"fmt"
	"github.com/VitalyDorozhkin/go-patterns/pkg/builder/computer"
	"github.com/VitalyDorozhkin/go-patterns/pkg/builder/computer-builder"
	"github.com/VitalyDorozhkin/go-patterns/pkg/builder/director"
)

func main() {
	appleDirector := director.NewDirector(computer_builder.NewAppleBuilder())
	xiaomiDirector := director.NewDirector(computer_builder.NewXiaomiBuilder())

	computers := []*computer.Computer{xiaomiDirector.ConstructPro(), xiaomiDirector.ConstructAir(), appleDirector.ConstructPro(), appleDirector.ConstructAir()}

	for i, v := range computers {
		fmt.Printf("Vendor code: %d\ndescription:\n%s\n", i, v.Show())
	}
}
