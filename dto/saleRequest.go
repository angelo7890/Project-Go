package dto

type SaleRequestDTO struct {
	UserId   int `json:"user_id"`
	SectorId int `json:"sector_id"`
}

func (r *SaleRequestDTO) Validade() error {
	if r.UserId <= 0 {
		return ErrParamIsRequired("user_id", "int")
	}
	if r.SectorId <= 0 {
		return ErrParamIsRequired("sector_id", "int")
	}
	return nil
}

type ResponseSectorDTO struct {
	Id       int `json:"id"`
	UserId   int `json:"user_id"`
	SectorId int `json:"sector_id"`
}
