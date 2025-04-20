package dto

type SectorCreateDTO struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	ShowID   int    `json:"show_id"`
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

type ReponseSectorDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	ShowID   int    `json:"show_id"`
}
