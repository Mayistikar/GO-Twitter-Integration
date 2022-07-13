package domain

import (
	"errors"
)

type Profile struct {
	uuid        string
	image       string
	title       string
	description string
}

func (p *Profile) UUID() string {
	return p.uuid
}

func (p *Profile) Image() string {
	return p.image
}

func (p *Profile) Title() string {
	return p.title
}

func (p *Profile) Description() string {
	return p.description
}

var ErrorProfileInvalidImage = errors.New("invalid profile image")
var ErrorProfileInvalidTitle = errors.New("invalid profile title")
var ErrorProfileInvalidDescription = errors.New("invalid profile description")
