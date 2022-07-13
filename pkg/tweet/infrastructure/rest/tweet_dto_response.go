package rest

import "portfolio/pkg/tweet/domain"

type TweetDTOResponse struct {
	ID   string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}

func ToDTOResponse(args ...domain.Tweet) []TweetDTOResponse {
	var tweetDTOs []TweetDTOResponse
	for _, tweet := range args {
		var dto TweetDTOResponse
		dto.ID = tweet.ID()
		dto.Text = tweet.Text()
		tweetDTOs = append(tweetDTOs, dto)
	}
	return tweetDTOs
}
