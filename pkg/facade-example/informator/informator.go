package informator

import "fmt"

type Informator interface {
	InformNameLength(name string, min int, max int) string
	InformLastNameLength(lastname string, min int, max int) string
	InformAge(age int, min int, max int) string
}

type informator struct {
	maxLengthPattern string
	minLengthPattern string
	maxNumberPattern string
	minNumberPattern string
}

func (i *informator) InformNameLength(name string, min int, max int) (info string) {
	info = "Name Length: "
	if len(name) > max {
		info += fmt.Sprintf(i.maxLengthPattern, name, max)
	} else if len(name) < min {
		info += fmt.Sprintf(i.minLengthPattern, name, min)
	} else {
		info += "ok"
	}
	return
}

func (i *informator) InformLastNameLength(lastname string, min int, max int) (info string) {
	info = "LastName Length: "
	if len(lastname) > max {
		info += fmt.Sprintf(i.maxLengthPattern, lastname, max)
	} else if len(lastname) < min {
		info += fmt.Sprintf(i.minLengthPattern, lastname, min)
	} else {
		info += "ok"
	}
	return
}

func (i *informator) InformAge(age int, min int, max int) (info string) {
	info = "Age: "
	if age > max {
		info += fmt.Sprintf(i.maxNumberPattern, age, max)
	} else if age < min {
		info += fmt.Sprintf(i.minNumberPattern, age, min)
	} else {
		info += "ok"
	}
	return
}

func NewInformator(maxLengthPattern string, minLengthPattern string, maxNumberPattern string, minNumberPattern string) Informator {
	return &informator{
		maxLengthPattern: maxLengthPattern,
		minLengthPattern: minLengthPattern,
		maxNumberPattern: maxNumberPattern,
		minNumberPattern: minNumberPattern,
	}
}
