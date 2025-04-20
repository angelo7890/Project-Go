package database

import (
	"database/sql"
	"fmt"
	"ingressos-api/configuration"
	"os"

	_ "github.com/lib/pq"
)

var (
	logger *configuration.Logger
)
var DB *sql.DB

func Connect() (*sql.DB, error) {

	logger := configuration.GetLogger("postgre")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	tempConnStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		host, port, user, password,
	)

	tempDB, err := sql.Open("postgres", tempConnStr)
	if err != nil {
		logger.Errorf("erro ao conectar no banco postgres: %v", err)
		return nil, err
	}
	defer tempDB.Close()

	var exists bool
	err = tempDB.QueryRow("SELECT EXISTS(SELECT FROM pg_database WHERE datname = $1)", dbname).Scan(&exists)
	if err != nil {
		logger.Errorf("erro ao verificar existência do banco: %v", err)
		return nil, err
	}

	if !exists {
		_, err = tempDB.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			logger.Errorf("erro ao criar banco de dados: %v", err)
			return nil, err
		}
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Errorf("erro ao conectar no banco de dados: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.Errorf("erro ao verificar conexão com o banco: %v", err)
		return nil, err
	}

	schema, err := os.ReadFile("database/sql/init_db.sql")
	if err != nil {
		logger.Errorf("erro ao ler arquivo init_schema.sql: %v", err)
		return nil, err
	}

	if _, err := db.Exec(string(schema)); err != nil {
		logger.Errorf("erro ao executar init_schema.sql: %v", err)
		return nil, err
	}

	logger.Info("Banco de dados conectado e tabelas prontas.")
	return db, nil
}
