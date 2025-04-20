package dto

import (
	"time"
)

type EventCreateDTO struct {
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Location string    `json:"location"`
}

func (r *EventCreateDTO) Validate() error {
	if r.Name == "" {
		return ErrParamIsRequired("name", "string")
	}
	if r.Date.IsZero() {
		return ErrParamIsRequired("date", "date")
	}
	if r.Location == "" {
		return ErrParamIsRequired("location", "string")
	}
	return nil
}

type EventResponseDTO struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Location string    `json:"location"`
}
