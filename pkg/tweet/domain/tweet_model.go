package domain

import (
	"errors"
)

var (
	ErrorTweetInvalidID   = errors.New("invalid tweet id")
	ErrorTweetInvalidText = errors.New("invalid text id")
)

type Tweet struct {
	id   string
	text string
}

func (t *Tweet) ID() string {
	return t.id
}

func (t *Tweet) Text() string {
	return t.text
}

func (t *Tweet) SetID(id string) error {
	if id == "" {
		return ErrorTweetInvalidID
	}
	t.id = id
	return nil
}

func (t *Tweet) SetText(text string) error {
	if text == "" {
		return ErrorTweetInvalidText
	}
	t.text = text
	return nil
}
