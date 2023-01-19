package stats

import (
	"encoding/json"
	"net/http"

	"github.com/luisrosas/examen-mercadolibre/internal/dna/domain"
)

type statsResponse struct {
	CountMutantDna uint    `json:"count_mutant_dna"`
	CountHumanDna  uint    `json:"count_human_dna"`
	Ratio          float32 `json:"ratio"`
}

type dnaStatisticsUserCase interface {
	Handle() (domain.DnaStatistics, error)
}

func Handle(useCase dnaStatisticsUserCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		stats, err := useCase.Handle()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		statsResponse := statsResponse{
			CountMutantDna: stats.Mutants,
			CountHumanDna:  stats.Humans,
			Ratio:          stats.Ratio(),
		}

		jsonResp, err := json.Marshal(statsResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}
}
