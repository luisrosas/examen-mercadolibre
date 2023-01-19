package application

import "github.com/stretchr/testify/mock"

// MockDnaValidator is a mock type for the DnaValidator type
type MockDnaValidator struct {
	mock.Mock
}

// IsMutant provides a mock function with given fields: dnaChain
func (_m *MockDnaValidator) IsMutant(dnaChain []string) bool {
	args := _m.Called(dnaChain)
	return args.Get(0).(bool)
}

// IsChainValid provides a mock function with given fields: dnaChain
func (_m *MockDnaValidator) IsChainValid(dnaChain []string) bool {
	args := _m.Called(dnaChain)
	return args.Get(0).(bool)
}
