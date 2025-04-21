package repository

import (
	"database/sql"
	"ingressos-api/dto"
)

func CreateSector(db *sql.DB, sector dto.SectorCreateDTO) error {
	query := `INSERT INTO setor (nome, capacidade, show_id) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, sector.Name, sector.Capacity, sector.ShowID)
	return err
}

func DeleteSector(db *sql.DB, id int) error {
	query := `DELETE FROM setor WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}

func GetSectorsByEventID(db *sql.DB, eventID int) ([]dto.ReponseSectorDTO, error) {
	query := `SELECT id, nome, capacidade, show_id FROM setor WHERE show_id = $1`
	rows, err := db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sectors []dto.ReponseSectorDTO
	for rows.Next() {
		var s dto.ReponseSectorDTO
		if err := rows.Scan(&s.Id, &s.Name, &s.Capacity, &s.ShowID); err != nil {
			return nil, err
		}
		sectors = append(sectors, s)
	}
	return sectors, nil
}
