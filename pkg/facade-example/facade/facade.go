package facade

import (
	"fmt"
)

type validator interface {
	CheckNameLength(name string) (bool, int, int)
	CheckLastNameLength(lastname string) (bool, int, int)
	CheckAge(age int) (bool, int, int)
}

type informator interface {
	InformNameLength(name string, min int, max int) string
	InformLastNameLength(lastname string, min int, max int) string
	InformAge(age int, min int, max int) string
}

type user = interface {
	Info() string
}

type factory interface {
	NewUser(name string, lastname string, age int) user
}

type UserCreator interface {
	NewUser(name string, lastname string, age int) (user, error)
}

type userCreator struct {
	validator  validator
	informator informator
	factory    factory
}

func (u *userCreator) NewUser(name string, lastname string, age int) (user user, err error) {
	if res, min, max := u.validator.CheckNameLength(name); !res {
		err = fmt.Errorf(u.informator.InformNameLength(name, min, max))
		return
	}
	if res, min, max := u.validator.CheckLastNameLength(name); !res {
		err = fmt.Errorf(u.informator.InformLastNameLength(lastname, min, max))
		return
	}
	if res, min, max := u.validator.CheckAge(age); !res {
		err = fmt.Errorf(u.informator.InformAge(age, min, max))
		return
	}
	user = u.factory.NewUser(name, lastname, age)
	return
}

func NewUserCreator(f factory, v validator, i informator) (u UserCreator) {
	u = &userCreator{
		validator:  v,
		informator: i,
		factory:    f,
	}
	return
}
