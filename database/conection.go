package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
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
		return nil, fmt.Errorf("erro ao conectar no banco postgres: %w", err)
	}
	defer tempDB.Close()

	var exists bool
	err = tempDB.QueryRow("SELECT EXISTS(SELECT FROM pg_database WHERE datname = $1)", dbname).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar existência do banco: %w", err)
	}

	if !exists {
		_, err = tempDB.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar banco de dados: %w", err)
		}
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar no banco de dados: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao verificar conexão com o banco: %w", err)
	}

	schema, err := os.ReadFile("database/sql/init_db.sql")
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo init_db.sql: %w", err)
	}

	if _, err := db.Exec(string(schema)); err != nil {
		return nil, fmt.Errorf("erro ao executar init_db.sql: %w", err)
	}

	fmt.Println("Banco de dados conectado e tabelas prontas.")
	return db, nil
}
