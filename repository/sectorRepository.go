package repository

import (
	"database/sql"
	"fmt"
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

func UpdateTicketSector(tx *sql.Tx, ticketID int, newSectorID int) (*dto.ResponseBuyTicketDTO, error) {
	var capacidade, totalVendas int

	err := tx.QueryRow("SELECT capacidade FROM setor WHERE id = $1 FOR UPDATE", newSectorID).Scan(&capacidade)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar capacidade do novo setor: %w", err)
	}

	err = tx.QueryRow("SELECT COUNT(*) FROM venda_ingresso WHERE setor_id = $1", newSectorID).Scan(&totalVendas)
	if err != nil {
		return nil, fmt.Errorf("erro ao contar vendas no novo setor: %w", err)
	}

	if totalVendas >= capacidade {
		return nil, fmt.Errorf("novo setor está esgotado")
	}

	_, err = tx.Exec("UPDATE venda_ingresso SET setor_id = $1 WHERE id = $2", newSectorID, ticketID)
	if err != nil {
		return nil, fmt.Errorf("erro ao atualizar setor do ingresso: %w", err)
	}

	var ticket dto.ResponseBuyTicketDTO
	err = tx.QueryRow(`SELECT vi.id AS ticket_id, vi.usuario_id AS user_id, vi.setor_id AS sector_id, s.show_id AS event_id 
		FROM venda_ingresso vi
		JOIN setor s ON vi.setor_id = s.id
		WHERE vi.id = $1`, ticketID).
		Scan(&ticket.TicketId, &ticket.UserId, &ticket.SectorId, &ticket.EventId)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar ingresso atualizado: %w", err)
	}

	return &ticket, nil
}
