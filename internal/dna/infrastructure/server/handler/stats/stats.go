package stats

import (
	"encoding/json"
	"net/http"

	"github.com/luisrosas/examen-mercadolibre/internal/dna/application"
)

type statsResponse struct {
	CountMutantDna uint    `json:"count_mutant_dna"`
	CountHumanDna  uint    `json:"count_human_dna"`
	Ratio          float32 `json:"ratio"`
}

func Handle(dnaStatisticsUserCase application.DnaStatisticsUserCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		stats, err := dnaStatisticsUserCase.Handle()
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
