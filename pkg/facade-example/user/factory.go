package user

// Factory ...
type Factory interface {
	NewUser(name string, lastname string, age int) User
}

type factory struct{}

// NewUser ...
func (f *factory) NewUser(name string, lastname string, age int) User {
	return &user{
		name:     name,
		lastname: lastname,
		age:      age,
	}
}

// NewFactory initializes the Factory
func NewFactory() Factory {
	return &factory{}
}
