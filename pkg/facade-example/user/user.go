package user

import "fmt"

type User = interface {
	Info() string
}

type user struct {
	name     string
	lastname string
	age      int
}

func (u *user) Info() string {
	return fmt.Sprintf("user: {\nname: %s,\nlastname: %s,\nage: %d\n}", u.name, u.lastname, u.age)
}
