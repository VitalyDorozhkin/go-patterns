package user

type Factory interface {
	NewUser(name string, lastname string, age int) User
}

type factory struct{}

func (f *factory) NewUser(name string, lastname string, age int) User {
	return &user{
		name:     name,
		lastname: lastname,
		age:      age,
	}
}

func NewFactory() Factory {
	return &factory{}
}
