package twitter

import (
	"errors"
	"fmt"
	"net/http"
	"portfolio/pkg/tweet/domain"
	"portfolio/shared/rest"
)

var (
	ErrorTwitterResponse = errors.New("invalid response, twitter unavailable")
)

type Repository struct {
	rest *rest.Client
}

func NewRepository(client *rest.Client) *Repository {
	return &Repository{client}
}

func (t *Repository) FindMany(max int64) ([]domain.Tweet, error) {
	response, err := t.rest.Client.R().
		SetQueryParam("max_results", fmt.Sprint(max)).
		Get(t.rest.Config.ThirdParty.Twitter.Host)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, ErrorTwitterResponse
	}

	tweets, err := MapTweets(response.Body())
	if err != nil {
		return nil, err
	}

	return tweets, nil
}
