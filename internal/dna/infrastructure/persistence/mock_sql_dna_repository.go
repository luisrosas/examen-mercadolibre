package persistence

import (
	"github.com/luisrosas/examen-mercadolibre/internal/dna/domain"

	"github.com/stretchr/testify/mock"
)

// MockSQLDnaRepository is a mock type for the SQLDnaRepository type
type MockSQLDnaRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: dna
func (_m *MockSQLDnaRepository) Save(dna domain.Dna) error {
	args := _m.Called(dna)
	return args.Error(0)
}

// Find provides a mock function with given fields: dnaChain
func (_m *MockSQLDnaRepository) Find(dnaChain []string) (domain.Dna, error) {
	args := _m.Called(dnaChain)
	return args.Get(0).(domain.Dna), args.Error(1)
}

// GetStatistics provides a mock function
func (_m *MockSQLDnaRepository) GetStatistics() (domain.DnaStatistics, error) {
	args := _m.Called()
	return args.Get(0).(domain.DnaStatistics), args.Error(1)
}
