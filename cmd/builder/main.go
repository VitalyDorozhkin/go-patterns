package main

import (
	. "../../pkg/builder"
	"fmt"
)

func main() {
	appleDirector := Director{new(AppleBuilder)}
	xiaomiDirector := Director{new(XiaomiBuilder)}

	computers := []*Computer{xiaomiDirector.ConstructPro(), xiaomiDirector.ConstructAir(), appleDirector.ConstructPro(), appleDirector.ConstructAir()}

	for i, v := range computers {
		fmt.Printf("Vendor code: %d\ndescription:\n%s\n", i, v.Show())
	}
}
