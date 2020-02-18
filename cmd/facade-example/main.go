package main

import (
	"fmt"

	"github.com/VitalyDorozhkin/go-patterns/pkg/facade-example/facade"
	"github.com/VitalyDorozhkin/go-patterns/pkg/facade-example/informator"
	"github.com/VitalyDorozhkin/go-patterns/pkg/facade-example/models"
	"github.com/VitalyDorozhkin/go-patterns/pkg/facade-example/user"
	"github.com/VitalyDorozhkin/go-patterns/pkg/facade-example/validator"
)

func main() {
	factory := user.NewFactory()
	i := informator.NewInformator(
		"%s is too long! max %d",
		"%s is too short! min %d",
		"%d is too big! max %d",
		"%d is too small! min %d",
	)
	v := validator.NewValidator(
		models.Interval{Min: 3, Max: 8},
		models.Interval{Min: 3, Max: 8},
		models.Interval{Min: 18, Max: 100},
	)
	creator := facade.NewUserCreator(factory, v, i)
	if user1, err1 := creator.NewUser("Ivan", "Ivanov", 18); err1 != nil {
		fmt.Printf("user1: %s\n", err1)
	} else {
		fmt.Println(user1.Info())
	}

	if user2, err2 := creator.NewUser("Vasya", "Vasykin", 8); err2 != nil {
		fmt.Printf("user2: %s\n", err2)
	} else {
		fmt.Println(user2.Info())
	}

	if user3, err3 := creator.NewUser("Petya", "Petrovich", 81); err3 != nil {
		fmt.Printf("user3: %s\n", err3)
	} else {
		fmt.Println(user3.Info())
	}
}
