package twitter

import (
	"encoding/json"
	"portfolio/pkg/tweet/domain"
)

type MetaDTOResponse struct {
	NextToken   string `json:"next_token,omitempty"`
	ResultCount int    `json:"result_count,omitempty"`
	NewestID    string `json:"newest_id,omitempty"`
	OldestID    string `json:"oldest_id,omitempty"`
}

type TweetDTOResponse struct {
	Data []TweetDTO      `json:"data,omitempty"`
	Meta MetaDTOResponse `json:"meta"`
}

type TweetDTO struct {
	ID   string `form:"id" json:"id"`
	Text string `form:"text" json:"text"`
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

func MapTweets(twitterResponse []byte) ([]domain.Tweet, error) {
	var tweetDTOResponse TweetDTOResponse
	if err := json.Unmarshal(twitterResponse, &tweetDTOResponse); err != nil {
		return nil, err
	}

	var models []domain.Tweet

	for _, tweet := range tweetDTOResponse.Data {
		var model domain.Tweet

		if err := model.SetID(tweet.ID); err != nil {
			return nil, err
		}

		if err := model.SetText(tweet.Text); err != nil {
			return nil, err
		}

		models = append(models, model)
	}
	return models, nil
}
