package client

import (
	"net/http"
	"time"
)

const (
	BaseCoinDataURL = "https://api.coingecko.com/api/v3/coins/"
)

type Client struct {
	Client   *http.Client
	geckoKey string
}

func NewClient(timeout time.Duration, geckoKey string) *Client {
	c := Client{
		Client: &http.Client{
			Timeout: timeout,
		},
		geckoKey: geckoKey,
	}
	return &c
}
