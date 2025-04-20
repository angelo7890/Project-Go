package dto

import (
	"errors"
	"time"
)

type EventCreateDTO struct {
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	Location         string    `json:"location"`
	AvailableTickets int       `json:"available_tickets"`
	TicketPrice      float64   `json:"ticket_price"`
}

func (r *EventCreateDTO) Validate() error {
	if r.Name == "" {
		return errors.New("nome do evento é obrigatório")
	}
	if r.Description == "" {
		return errors.New("descrição do evento é obrigatória")
	}
	if r.StartTime.IsZero() {
		return errors.New("data de início é obrigatória")
	}
	if r.EndTime.IsZero() {
		return errors.New("data de fim é obrigatória")
	}
	if r.Location == "" {
		return errors.New("localização do evento é obrigatória")
	}
	if r.AvailableTickets <= 0 {
		return errors.New("número de ingressos deve ser maior que zero")
	}
	if r.TicketPrice <= 0 {
		return errors.New("o valor do ingresso nao pode ser negativo")
	}
	if r.EndTime.Before(r.StartTime) {
		return errors.New("data de fim não pode ser antes da data de início")
	}
	return nil
}

type EventResponseDTO struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	Location         string    `json:"location"`
	AvailableTickets int       `json:"available_tickets"`
	TicketPrice      float64   `json:"ticket_price"`
}
