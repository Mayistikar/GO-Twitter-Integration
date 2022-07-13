package rest

import "errors"

var (
	ErrorProfileDTOInvalidUUID        = errors.New("invalid profile uuid")
	ErrorProfileDTOInvalidImage       = errors.New("invalid profile image")
	ErrorProfileDTOInvalidTitle       = errors.New("invalid profile title")
	ErrorProfileDTOInvalidDescription = errors.New("invalid profile description")
)
