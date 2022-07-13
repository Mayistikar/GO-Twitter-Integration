package rest

import (
	"errors"
	"portfolio/pkg/profile/domain"
)

var (
	ErrorProfileDTOInvalidUUID        = errors.New("invalid profile uuid")
	ErrorProfileDTOInvalidImage       = errors.New("invalid profile image")
	ErrorProfileDTOInvalidTitle       = errors.New("invalid profile title")
	ErrorProfileDTOInvalidDescription = errors.New("invalid profile description")
)

type ProfileDTO struct {
	UUID        string `json:"uuid" uri:"uuid" binding:"required,uuid"`
	Image       string `json:"image" binding:"omitempty"`
	Title       string `json:"title" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty"`
}

func (dto *ProfileDTO) SetUUID(UUID string) error {
	if UUID == "" {
		return ErrorProfileDTOInvalidUUID
	}
	dto.UUID = UUID
	return nil
}

func (dto *ProfileDTO) SetImage(image string) error {
	if image == "" {
		return ErrorProfileDTOInvalidImage
	}
	dto.Image = image
	return nil
}

func (dto *ProfileDTO) SetTitle(title string) error {
	if title == "" {
		return ErrorProfileDTOInvalidTitle
	}
	dto.Title = title
	return nil
}

func (dto *ProfileDTO) SetDescription(description string) error {
	if description == "" {
		return ErrorProfileDTOInvalidDescription
	}
	dto.Description = description
	return nil
}

func (dto *ProfileDTO) toProfile() (domain.Profile, error) {
	return domain.NewProfileBuilder().
		SetUUID(dto.UUID).
		SetImage(dto.Image).
		SetDescription(dto.Description).
		SetTitle(dto.Title).
		Build()
}

func (dto *ProfileDTO) mapProfile(profile domain.Profile) error {
	if err := dto.SetUUID(profile.UUID()); err != nil {
		return err
	}
	if err := dto.SetImage(profile.Image()); err != nil {
		return err
	}
	if err := dto.SetTitle(profile.Title()); err != nil {
		return err
	}
	if err := dto.SetDescription(profile.Description()); err != nil {
		return err
	}
	return nil
}
