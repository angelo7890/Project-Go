package repository

import (
	"database/sql"
	"ingressos-api/dto"
)

func CreateEvent(db *sql.DB, event dto.EventCreateDTO) error {
	query := `INSERT INTO show (nome, data, local) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, event.Name, event.Date, event.Location)
	return err
}

func GetAllEvents(db *sql.DB) ([]dto.EventResponseDTO, error) {
	rows, err := db.Query("SELECT id, nome, data, local FROM show")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []dto.EventResponseDTO
	for rows.Next() {
		var e dto.EventResponseDTO
		if err := rows.Scan(&e.Id, &e.Name, &e.Date, &e.Location); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func GetEventByID(db *sql.DB, id int) (*dto.EventResponseDTO, error) {
	var e dto.EventResponseDTO
	query := "SELECT id, nome, data, local FROM show WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&e.Id, &e.Name, &e.Date, &e.Location)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
