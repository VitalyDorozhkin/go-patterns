package main

import (
	"fmt"
	"github.com/VitalyDorozhkin/go-patterns/pkg/builder/computer"
	"github.com/VitalyDorozhkin/go-patterns/pkg/builder/concrete-builder/apple"
	"github.com/VitalyDorozhkin/go-patterns/pkg/builder/concrete-builder/xiaomi"
	"github.com/VitalyDorozhkin/go-patterns/pkg/builder/director"
)

func main() {
	appleDirector := director.NewService(apple.NewBuilder())
	xiaomiDirector := director.NewService(xiaomi.NewBuilder())

	computers := []*computer.Computer {
		xiaomiDirector.ConstructPro(),
		xiaomiDirector.ConstructAir(),
		appleDirector.ConstructPro(),
		appleDirector.ConstructAir(),
	}

	for i, v := range computers {
		fmt.Printf("Vendor code: %d\ndescription:\n%s\n", i, v.Show())
	}
}
