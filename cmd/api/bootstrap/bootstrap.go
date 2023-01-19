package bootstrap

import (
	"errors"
	"os"
	"strconv"

	"github.com/luisrosas/examen-mercadolibre/internal/dna/application"
	"github.com/luisrosas/examen-mercadolibre/internal/dna/infrastructure/persistence"
	"github.com/luisrosas/examen-mercadolibre/internal/dna/infrastructure/server"
)

func Initialize() (*server.Server, error) {
	dbConfig, err := initDBConfig()
	if err != nil {
		return nil, err
	}

	serverConfig, err := initServerConfig()
	if err != nil {
		return nil, err
	}

	repo, err := persistence.NewDnaRepository(dbConfig)
	if err != nil {
		return nil, err
	}

	validator := application.NewDnaValidator()
	dnaUseCase := application.NewDnaUseCase(repo, validator)
	dnaStatsUseCase := application.NewDnaStatsUserCase(repo)

	return server.NewServer(serverConfig, *dnaUseCase, *dnaStatsUseCase), nil
}

func initDBConfig() (persistence.SQLConfig, error) {
	config := persistence.SQLConfig{}

	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		return config, errors.New("not found DB_USER")
	}

	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return config, errors.New("not found DB_PASSWORD")
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return config, errors.New("not found DB_HOST")
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return config, errors.New("not found DB_PORT")
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return config, errors.New("not found DB_NAME")
	}

	port, err := strconv.Atoi(dbPort)
	if err != nil {
		return config, err
	}

	return persistence.SQLConfig{
		User:     dbUser,
		Password: dbPassword,
		Host:     dbHost,
		Port:     uint(port),
		Database: dbName,
	}, nil
}

func initServerConfig() (server.Config, error) {
	serverPort, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		return server.Config{}, errors.New("not found SERVER_PORT")
	}

	port, err := strconv.Atoi(serverPort)
	if err != nil {
		return server.Config{}, err
	}

	return server.Config{
		Port: uint(port),
	}, nil
}
