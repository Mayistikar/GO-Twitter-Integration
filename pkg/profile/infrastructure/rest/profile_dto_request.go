package rest

import (
	"portfolio/pkg/profile/domain"
)

type RequestDTO struct {
	UUID        string `uri:"uuid" binding:"required,uuid"`
	Image       string `json:"image" binding:"omitempty"`
	Title       string `json:"title" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty"`
}

func (dto *RequestDTO) toProfile() (domain.Profile, error) {
	return domain.NewProfileBuilder().
		SetUUID(dto.UUID).
		SetImage(dto.Image).
		SetDescription(dto.Description).
		SetTitle(dto.Title).
		Build()
}
