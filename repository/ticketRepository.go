package repository

import (
	"database/sql"
	"fmt"
	"ingressos-api/dto"
)

func BuyTicket(tx *sql.Tx, request dto.BuyTicketRequestDTO, db *sql.DB) (*dto.ResponseBuyTicketDTO, error) {
	// Bloquear o setor para garantir que ninguém mais altere enquanto verificamos a disponibilidade
	var setorId int
	var capacidade int
	err := tx.QueryRow("SELECT id, capacidade FROM setores WHERE id = $1 FOR UPDATE", request.SectorId).Scan(&setorId, &capacidade)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	// Verificar se há ingressos disponíveis no setor
	var ingressosDisponiveis int
	err = tx.QueryRow("SELECT COUNT(*) FROM ingressos WHERE setor_id = $1 AND status = 'disponível'", setorId).Scan(&ingressosDisponiveis)
	if err != nil {
		return nil, err
	}

	if ingressosDisponiveis <= 0 {
		return nil, err
	}

	// Marcar o ingresso como vendido
	var ingressoId int
	err = tx.QueryRow("UPDATE ingressos SET status = 'vendido' WHERE setor_id = $1 AND status = 'disponível' RETURNING id", setorId).Scan(&ingressoId)
	if err != nil {
		return nil, err
	}

	// Registrar a venda
	var vendaId int
	err = tx.QueryRow("INSERT INTO vendas (usuario_id, ingresso_id) VALUES ($1, $2) RETURNING id", request.UserId, ingressoId).Scan(&vendaId)
	if err != nil {
		return nil, err
	}

	// Recuperar event_id do ingresso
	var eventId int
	err = tx.QueryRow("SELECT evento_id FROM setores WHERE id = $1", request.SectorId).Scan(&eventId)
	if err != nil {
		return nil, err
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
