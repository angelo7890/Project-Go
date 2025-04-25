package repository

import (
	"database/sql"
	"fmt"
	"ingressos-api/dto"
)

func BuyTicket(tx *sql.Tx, request dto.BuyTicketRequestDTO) (*dto.ResponseBuyTicketDTO, error) {
	var capacidade, totalVendas int

	err := tx.QueryRow("SELECT capacidade FROM setor WHERE id = $1 FOR UPDATE", request.SectorId).Scan(&capacidade)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar capacidade: %w", err)
	}

	err = tx.QueryRow("SELECT COUNT(*) FROM venda_ingresso WHERE setor_id = $1", request.SectorId).Scan(&totalVendas)
	if err != nil {
		return nil, fmt.Errorf("erro ao contar vendas: %w", err)
	}

	if totalVendas >= capacidade {
		return nil, fmt.Errorf("setor esgotado")
	}

	var saleID int
	err = tx.QueryRow(
		"INSERT INTO venda_ingresso (usuario_id, setor_id) VALUES ($1, $2) RETURNING id",
		request.UserId, request.SectorId,
	).Scan(&saleID)
	if err != nil {
		return nil, fmt.Errorf("erro ao inserir venda: %w", err)
	}

	var eventID int
	err = tx.QueryRow("SELECT show_id FROM setor WHERE id = $1", request.SectorId).Scan(&eventID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar evento: %w", err)
	}

	ticket := &dto.ResponseBuyTicketDTO{
		SaleId:   saleID,
		UserId:   request.UserId,
		TicketId: saleID,
		SectorId: request.SectorId,
		EventId:  eventID,
	}

	return ticket, nil
}

func GetAllTicketsSoldRepository(db *sql.DB) ([]dto.ResponseBuyTicketDTO, error) {
	query := `
		SELECT v.id, v.usuario_id, i.id AS ticket_id, s.id AS sector_id, s.show_id AS event_id
		FROM venda_ingresso v
		JOIN ingresso i ON v.ingresso_id = i.id
		JOIN setor s ON i.setor_id = s.id
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar ingressos vendidos: %w", err)
	}
	defer rows.Close()

	var tickets []dto.ResponseBuyTicketDTO
	for rows.Next() {
		var t dto.ResponseBuyTicketDTO
		if err := rows.Scan(&t.SaleId, &t.UserId, &t.TicketId, &t.SectorId, &t.EventId); err != nil {
			return nil, fmt.Errorf("erro ao ler dados de venda: %w", err)
		}
		tickets = append(tickets, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre as linhas: %w", err)
	}
	return tickets, nil
}

func GetAllTicketsSoldByEventIDRepository(db *sql.DB, eventID int) ([]dto.ResponseBuyTicketDTO, error) {
	query := `
		SELECT v.id, v.usuario_id, i.id AS ticket_id, s.id AS sector_id, s.show_id AS event_id
		FROM venda_ingresso v
		JOIN ingresso i ON v.ingresso_id = i.id
		JOIN setor s ON i.setor_id = s.id
		WHERE s.show_id = $1
	`
	rows, err := db.Query(query, eventID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar ingressos vendidos para o evento: %w", err)
	}
	defer rows.Close()

	var tickets []dto.ResponseBuyTicketDTO
	for rows.Next() {
		var t dto.ResponseBuyTicketDTO
		if err := rows.Scan(&t.SaleId, &t.UserId, &t.TicketId, &t.SectorId, &t.EventId); err != nil {
			return nil, fmt.Errorf("erro ao ler dados de venda do evento: %w", err)
		}
		tickets = append(tickets, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre as linhas: %w", err)
	}
	return tickets, nil
}
