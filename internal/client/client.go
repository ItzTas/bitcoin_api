package client

import (
	"net/http"
	"time"
)

type Client struct {
	Client *http.Client
}

func NewClient(timeout time.Duration) *Client {
	c := Client{
		Client: &http.Client{
			Timeout: timeout,
		},
	}
	return &c
}
