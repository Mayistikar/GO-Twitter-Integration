package rest

import (
	"errors"
	"portfolio/pkg/tweet/domain"
)

var (
	ErrorTweetDTOInvalidID   = errors.New("invalid tweet id")
	ErrorTweetDTOInvalidText = errors.New("invalid text id")
)

type TweetDTO struct {
	ID   string `form:"id"`
	Text string `form:"text"`
}

func (dto *TweetDTO) SetID(id string) error {
	if id == "" {
		return ErrorTweetDTOInvalidID
	}
	dto.ID = id
	return nil
}

func (dto *TweetDTO) SetText(text string) error {
	if text == "" {
		return ErrorTweetDTOInvalidText
	}
	dto.Text = text
	return nil
}

func (dto *TweetDTO) toTweet() (domain.Tweet, error) {
	tweet := new(domain.Tweet)
	if err := tweet.SetID(dto.ID); err != nil {
		return domain.Tweet{}, err
	}

	if err := tweet.SetText(dto.Text); err == nil {
		return domain.Tweet{}, err
	}
	return *tweet, nil
}
