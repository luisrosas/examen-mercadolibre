package application

import (
	"strings"
)

type DnaValidator interface {
	IsMutant(dnaChain []string) bool
	IsChainValid(dnaChain []string) bool
}

type dnaValidator struct {
}

const (
	coreBaseA = "A"
	coreBaseT = "T"
	coreBaseC = "C"
	coreBaseG = "G"
)

func NewDnaValidator() DnaValidator {
	return dnaValidator{}
}

func (d dnaValidator) IsMutant(dnaChain []string) bool {
	for i := 0; i < len(dnaChain); i++ {
		for j := 0; j < len(dnaChain[i]); j++ {
			// vertical
			if i <= len(dnaChain)-4 {
				sequence := string(
					[]byte{
						dnaChain[i][j],
						dnaChain[i+1][j],
						dnaChain[i+2][j],
						dnaChain[i+3][j],
					},
				)
				if isSequenceEqual(sequence) {
					return true
				}
			}

			// horizontal
			if j <= len(dnaChain[i])-4 {
				sequence := dnaChain[i][j : j+4]
				if isSequenceEqual(sequence) {
					return true
				}
			}

			// oblique
			if i <= len(dnaChain)-4 {
				var sequence string

				if j <= len(dnaChain[i])-4 {
					sequence = string(
						[]byte{
							dnaChain[i][j],
							dnaChain[i+1][j+1],
							dnaChain[i+2][j+2],
							dnaChain[i+3][j+3],
						},
					)
				}

				if j >= 4 {
					sequence = string(
						[]byte{
							dnaChain[i][j],
							dnaChain[i+1][j-1],
							dnaChain[i+2][j-2],
							dnaChain[i+3][j-3],
						},
					)
				}

				if sequence != "" && isSequenceEqual(sequence) {
					return true
				}
			}
		}
	}

	return false
}

func isSequenceEqual(sequence string) bool {
	for i := 1; i < len(sequence); i++ {
		if sequence[0] != sequence[i] {
			return false
		}
	}

	return true
}

func (d dnaValidator) IsChainValid(dnaChain []string) bool {
	if !validateSize(dnaChain) || !validateContent(dnaChain) {
		return false
	}

	return true
}
func validateSize(dnaChain []string) bool {
	lenChain := len(dnaChain)
	if lenChain >= 4 {
		for _, dnaSec := range dnaChain {
			if len(dnaSec) != lenChain {
				return false
			}
		}

		return true
	}

	return false
}

func validateContent(dnaChain []string) bool {
	var counter int

	dnaChainComplete := strings.Join(dnaChain, "")

	counter = strings.Count(dnaChainComplete, coreBaseA)
	counter += strings.Count(dnaChainComplete, coreBaseT)
	counter += strings.Count(dnaChainComplete, coreBaseC)
	counter += strings.Count(dnaChainComplete, coreBaseG)

	return counter == len(dnaChainComplete)
}
