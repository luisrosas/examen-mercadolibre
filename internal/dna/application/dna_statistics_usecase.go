package application

import "github.com/luisrosas/examen-mercadolibre/internal/dna/domain"

type DnaStatisticsUserCase struct {
	dnaRepository domain.DnaRepository
}

func NewDnaStatsUserCase(repo domain.DnaRepository) *DnaStatisticsUserCase {
	return &DnaStatisticsUserCase{
		dnaRepository: repo,
	}
}

func (d DnaStatisticsUserCase) Handle() (domain.DnaStatistics, error) {
	return d.dnaRepository.GetStatistics()
}
