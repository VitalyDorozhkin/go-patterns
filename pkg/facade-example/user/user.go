package user

import "fmt"

// User ...
type User = interface {
	Info() string
}

type user struct {
	name     string
	lastname string
	age      int
}

// Info return user information
func (u *user) Info() string {
	return fmt.Sprintf("user: {\nname: %s,\nlastname: %s,\nage: %d\n}", u.name, u.lastname, u.age)
}
