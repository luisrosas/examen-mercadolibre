package domain

import (
	"errors"
)

type Dna struct {
	Chain  []string
	Mutant bool
}

type DnaStatistics struct {
	Humans  uint
	Mutants uint
}

type DnaRepository interface {
	Save(dna Dna) error
	Find(dnaChain []string) (Dna, error)
	GetStatistics() (DnaStatistics, error)
}

var ErrNotFound = errors.New("not found")

func (d *Dna) IsMutant() bool {
	return d.Mutant
}

func (d *DnaStatistics) Ratio() float32 {
	if d.Humans > 0 {
		return float32(d.Mutants) / float32(d.Humans)
	}

	return float32(d.Mutants)
}
