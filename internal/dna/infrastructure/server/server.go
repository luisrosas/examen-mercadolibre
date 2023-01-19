package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luisrosas/examen-mercadolibre/internal/dna/application"
	"github.com/luisrosas/examen-mercadolibre/internal/dna/infrastructure/server/handler/mutant"
	"github.com/luisrosas/examen-mercadolibre/internal/dna/infrastructure/server/handler/stats"

	"github.com/gorilla/mux"
)

type Config struct {
	Port uint
}

type Server struct {
	Port                  uint
	Engine                *http.Server
	DnaUseCase            application.DnaUseCase
	DnaStatisticsUserCase application.DnaStatisticsUserCase
}

func NewServer(
	config Config,
	dnaUseCase application.DnaUseCase,
	dnaStatisticsUserCase application.DnaStatisticsUserCase,
) *Server {
	return &Server{
		Port:                  config.Port,
		DnaUseCase:            dnaUseCase,
		DnaStatisticsUserCase: dnaStatisticsUserCase,
	}
}

func (s *Server) Run() error {
	s.Engine = &http.Server{
		Addr:    fmt.Sprintf(":%v", s.Port),
		Handler: s.registerRoutes(),
	}
	log.Printf("Server started in port %d\n", s.Port)

	return s.Engine.ListenAndServe()
}

func (s *Server) registerRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/mutant", mutant.Handle(s.DnaUseCase)).Methods(http.MethodPost)
	r.HandleFunc("/stats", stats.Handle(s.DnaStatisticsUserCase)).Methods(http.MethodGet)

	return r
}
