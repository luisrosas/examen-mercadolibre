package application

import (
	"github.com/luisrosas/examen-mercadolibre/internal/dna/domain"

	"github.com/stretchr/testify/mock"
)

// MockDnaUseCase is a mock type for the DnaUseCase type
type MockDnaUseCase struct {
	mock.Mock
}

// Handle provides a mock function with given fields: dnaChain
func (_m *MockDnaUseCase) Handle(dnaChain []string) (domain.Dna, error) {
	args := _m.Called(dnaChain)
	return args.Get(0).(domain.Dna), args.Error(1)
}
