package mutant

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/luisrosas/examen-mercadolibre/internal/dna/application"
)

var (
	mutantResponse = http.StatusOK
	humanResponse  = http.StatusForbidden
)

type dnaRequest struct {
	Dna []string `json:"dna"`
}

func Handle(service application.DnaUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		var dna dnaRequest
		err = json.Unmarshal(b, &dna)
		if err != nil {
			log.Printf("Json unmarshal: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		responseDna, err := service.Handle(dna.Dna)
		if err != nil {
			log.Println(err)
			switch {
			case errors.Is(err, application.ErrInvalidChain):
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			return
		}

		if responseDna.IsMutant() {
			w.WriteHeader(mutantResponse)
		} else {
			w.WriteHeader(humanResponse)
		}
	}
}
