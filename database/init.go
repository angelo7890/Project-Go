package database

import (
	"database/sql"
	"fmt"
)

var (
	db *sql.DB
)

func Initialize() error {
	var err error
	db, err = Connect()

	if err != nil {
		return fmt.Errorf("error initializing postgre: %v", err)
	}
	return nil
}

func GetDB() *sql.DB {
	return db
}
