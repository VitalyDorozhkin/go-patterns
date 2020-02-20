package validator

import "github.com/stretchr/testify/mock"

// Mock of validator
type Mock struct {
	mock.Mock
}

// CheckNameLength ...
func (m *Mock) CheckNameLength(name string) (bool, int, int) {
	args := m.Called(name)
	return args.Bool(0), args.Int(1), args.Int(2)
}

// CheckLastNameLength ...
func (m *Mock) CheckLastNameLength(name string) (bool, int, int) {
	args := m.Called(name)
	return args.Bool(0), args.Int(1), args.Int(2)
}

// CheckAge ...
func (m *Mock) CheckAge(age int) (bool, int, int) {
	args := m.Called(age)
	return args.Bool(0), args.Int(1), args.Int(2)
}
