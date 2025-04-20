package main

import (
	"ingressos-api/configuration"
	"ingressos-api/database"
	"ingressos-api/router"
	"log"

	"github.com/joho/godotenv"
)

var (
	logger *configuration.Logger
)

func main() {
	logger = configuration.GetLogger("main")

	error := godotenv.Load()
	if error != nil {
		log.Fatal("Erro ao carregar .env")
	}

	err := database.Initialize()
	if err != nil {
		logger.Errorf("Error ao inicar banco de dados: %v", err)
		return
	}

	router.Initialize()

}
