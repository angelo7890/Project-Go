package repository

import (
	"database/sql"
	"fmt"
	"ingressos-api/dto"
)

func BuyTicket(transaction *sql.Tx, request dto.BuyTicketRequestDTO) (*dto.ResponseBuyTicketDTO, error) {
	// Bloqueia o setor durante a transação
	var capacidade, ingressosVendidos int
	err := transaction.QueryRow("SELECT capacidade FROM setores WHERE id = $1 FOR UPDATE", request.SectorId).Scan(&capacidade)
	if err != nil {
		return nil, fmt.Errorf("setor não encontrado: %v", err)
	}

	// Verifica quantos ingressos já foram vendidos
	err = transaction.QueryRow("SELECT COUNT(*) FROM ingressos WHERE setor_id = $1 AND status = 'vendido'", request.SectorId).Scan(&ingressosVendidos)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar ingressos vendidos: %v", err)
	}

	if ingressosVendidos >= capacidade {
		return nil, fmt.Errorf("ingressos esgotados para este setor")
	}

	// Insere o ingresso como vendido
	var ingressoId int
	err = transaction.QueryRow("INSERT INTO ingressos (setor_id, status) VALUES ($1, 'vendido') RETURNING id", request.SectorId).Scan(&ingressoId)
	if err != nil {
		return nil, fmt.Errorf("erro ao registrar ingresso: %v", err)
	}

	// Registra a venda
	var vendaId int
	err = transaction.QueryRow("INSERT INTO vendas (usuario_id, ingresso_id) VALUES ($1, $2) RETURNING id", request.UserId, ingressoId).Scan(&vendaId)
	if err != nil {
		return nil, fmt.Errorf("erro ao registrar venda: %v", err)
	}

	// Busca o evento do setor
	var eventId int
	err = transaction.QueryRow("SELECT show_id FROM setores WHERE id = $1", request.SectorId).Scan(&eventId)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar evento: %v", err)
	}

	ticket := &dto.ResponseBuyTicketDTO{
		SaleId:   vendaId,
		UserId:   request.UserId,
		TicketId: ingressoId,
		SectorId: request.SectorId,
		EventId:  eventId,
	}

	return ticket, nil
}

func GetAllTicketsSoldRepository(db *sql.DB) ([]dto.ResponseBuyTicketDTO, error) {
	query := `
		SELECT v.id, v.usuario_id, i.id, s.id, s.evento_id
		FROM vendas v
		JOIN ingressos i ON v.ingresso_id = i.id
		JOIN setores s ON i.setor_id = s.id
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

	return tickets, nil
}

func GetAllTicketsSoldByEventIDRepository(db *sql.DB, eventID int) ([]dto.ResponseBuyTicketDTO, error) {
	query := `
		SELECT v.id, v.usuario_id, i.id, s.id, s.evento_id
		FROM vendas v
		JOIN ingressos i ON v.ingresso_id = i.id
		JOIN setores s ON i.setor_id = s.id
		WHERE s.evento_id = $1
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

	return tickets, nil
}
