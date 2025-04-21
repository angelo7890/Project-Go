package dto

type BuyTicketRequestDTO struct {
	UserId   int `json:"user_id"`
	SectorId int `json:"sector_id"`
}

func (r *BuyTicketRequestDTO) Validade() error {
	if r.UserId <= 0 {
		return ErrParamIsRequired("user_id", "int")
	}
	if r.SectorId <= 0 {
		return ErrParamIsRequired("sector_id", "int")
	}
	return nil
}

type ResponseBuyTicketDTO struct {
	SaleId   int `json:"sale_id"`
	UserId   int `json:"user_id"`
	TicketId int `json:"ticket_id"`
	SectorId int `json:"sector_id"`
	EventId  int `json:"event_id"`
}
