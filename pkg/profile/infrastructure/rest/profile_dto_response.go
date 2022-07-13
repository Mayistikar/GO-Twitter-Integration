package rest

import "portfolio/pkg/profile/domain"

type ResponseDTO struct {
	UUID        string `json:"uuid" binding:"required,uuid"`
	Image       string `json:"image" binding:"omitempty"`
	Title       string `json:"title" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty"`
}

func (dto *ResponseDTO) SetUUID(UUID string) error {
	if UUID == "" {
		return ErrorProfileDTOInvalidUUID
	}
	dto.UUID = UUID
	return nil
}

func (dto *ResponseDTO) SetImage(image string) error {
	if image == "" {
		return ErrorProfileDTOInvalidImage
	}
	dto.Image = image
	return nil
}

func (dto *ResponseDTO) SetTitle(title string) error {
	if title == "" {
		return ErrorProfileDTOInvalidTitle
	}
	dto.Title = title
	return nil
}

func (dto *ResponseDTO) SetDescription(description string) error {
	if description == "" {
		return ErrorProfileDTOInvalidDescription
	}
	dto.Description = description
	return nil
}

func (dto *ResponseDTO) mapProfile(profile domain.Profile) error {
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
