package informator

import "github.com/stretchr/testify/mock"

// Mock of validator
type Mock struct {
	mock.Mock
}

// InformNameLength ...
func (m *Mock) InformNameLength(name string, min int, max int) string {
	args := m.Called(name, min, max)
	return args.String(0)
}

// InformLastNameLength ...
func (m *Mock) InformLastNameLength(lastname string, min int, max int) string {
	args := m.Called(lastname, min, max)
	return args.String(0)
}

// InformAge ...
func (m *Mock) InformAge(age int, min int, max int) string {
	args := m.Called(age, min, max)
	return args.String(0)
}
