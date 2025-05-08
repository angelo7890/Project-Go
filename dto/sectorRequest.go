package dto

type SectorCreateDTO struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	ShowID   int    `json:"show_id"`
}
type UpdateSectorDTO struct {
	TicketId    int `json:"ticket_id"`
	NewSectorId int `json:"new_sector_id"`
}

func (r *SectorCreateDTO) Validade() error {
	if r.Name == "" {
		return ErrParamIsRequired("name", "string")
	}
	if r.Capacity <= 0 {
		return ErrParamIsRequired("capacity", "int")
	}
	if r.ShowID <= 0 {
		return ErrParamIsRequired("show_id", "int")
	}
	return nil
}

func (r *UpdateSectorDTO) Validade() error {
	if r.TicketId <= 0 {
		return ErrParamIsRequired("ticket_id", "int")
	}
	if r.NewSectorId <= 0 {
		return ErrParamIsRequired("new_sector_id", "int")
	}
	return nil
}

type ReponseSectorDTO struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	ShowID   int    `json:"show_id"`
}
