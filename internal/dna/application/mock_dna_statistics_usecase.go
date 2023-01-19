package application

import (
	"github.com/luisrosas/examen-mercadolibre/internal/dna/domain"

	"github.com/stretchr/testify/mock"
)

// MockDnaStatisticsUserCase is a mock type for the DnaStatisticsUserCase type
type MockDnaStatisticsUserCase struct {
	mock.Mock
}

// Handle provides a mock function
func (_m *MockDnaStatisticsUserCase) Handle() (domain.DnaStatistics, error) {
	args := _m.Called()
	return args.Get(0).(domain.DnaStatistics), args.Error(1)
}
