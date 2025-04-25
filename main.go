package main

import (
	"ingressos-api/database"
	"ingressos-api/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	if err := database.Initialize(); err != nil {
		log.Fatalf("Erro ao iniciar banco de dados: %v", err)
	}

	router.Initialize()
}
