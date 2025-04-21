package handler

import (
	"database/sql"
	"ingressos-api/configuration"
	"ingressos-api/database"
)

var (
	logger *configuration.Logger
	db     *sql.DB
)

func InitializeHandler() {
	logger = configuration.GetLogger("handler")
	db = database.GetDB()
}
