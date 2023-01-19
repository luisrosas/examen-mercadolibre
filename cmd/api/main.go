package main

import (
	"log"

	"github.com/luisrosas/examen-mercadolibre/cmd/api/bootstrap"
)

func main() {
	server, err := bootstrap.Initialize()
	if err != nil {
		panic(err)
	}

	log.Fatal(server.Run())
}
