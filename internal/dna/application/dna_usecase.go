package application

import (
	"errors"

	"github.com/luisrosas/examen-mercadolibre/internal/dna/domain"
)

type DnaUseCase struct {
	dnaRepository domain.DnaRepository
	dnaValidator  DnaValidator
}

var (
	ErrUnexpected   = errors.New("unexpected error")
	ErrInvalidChain = errors.New("invalid chain dna")
)

func NewDnaUseCase(repo domain.DnaRepository, validator DnaValidator) *DnaUseCase {
	return &DnaUseCase{
		dnaRepository: repo,
		dnaValidator:  validator,
	}
}

func (d DnaUseCase) Handle(dnaChain []string) (domain.Dna, error) {
	ok := d.dnaValidator.IsChainValid(dnaChain)
	if !ok {
		return domain.Dna{}, ErrInvalidChain
	}

	dnaFound, err := d.dnaRepository.Find(dnaChain)
	if err == nil {
		return dnaFound, nil
	}

	if !errors.Is(err, domain.ErrNotFound) {
		return domain.Dna{}, ErrUnexpected
	}

	dna := domain.Dna{
		Chain:  dnaChain,
		Mutant: d.dnaValidator.IsMutant(dnaChain),
	}
	err = d.dnaRepository.Save(dna)
	if err != nil {
		return domain.Dna{}, ErrUnexpected
	}

	return dna, nil
}
