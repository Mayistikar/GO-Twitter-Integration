package rest

import (
	"github.com/go-resty/resty/v2"
	rest "portfolio/shared/config"
)

type Client struct {
	Config *rest.Config
	Client *resty.Client
}

func NewClient(config *rest.Config) *Client {
	client := resty.New()
	client.SetAuthToken(config.ThirdParty.Twitter.AuthToken)
	client.SetHeader("Accept", "application/json")
	return &Client{
		config,
		client,
	}
}
