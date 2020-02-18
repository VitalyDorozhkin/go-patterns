package validator

import "github.com/VitalyDorozhkin/go-patterns/pkg/facade-example/models"

type Validator interface {
	CheckNameLength(name string) (bool, int, int)
	CheckLastNameLength(lastname string) (bool, int, int)
	CheckAge(age int) (bool, int, int)
}

type validator struct {
	nameLength     models.Interval
	lastnameLength models.Interval
	age            models.Interval
}

func (v *validator) CheckNameLength(name string) (bool, int, int) {
	return v.nameLength.Include(len(name)), v.nameLength.Min, v.nameLength.Max
}

func (v *validator) CheckLastNameLength(lastname string) (bool, int, int) {
	return v.lastnameLength.Include(len(lastname)), v.nameLength.Min, v.nameLength.Max
}

func (v *validator) CheckAge(age int) (bool, int, int) {
	return v.age.Include(age), v.age.Min, v.age.Max
}

func NewValidator(nameLength models.Interval, lastnameLength models.Interval, age models.Interval) Validator {
	return &validator{
		nameLength:     nameLength,
		lastnameLength: lastnameLength,
		age:            age,
	}
}
