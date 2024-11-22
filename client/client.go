package client

import (
	"fmt"
	"net/http"
)

type Client struct {
	client *http.Client
	host   string
	UserID string
}

func NewClient(host string, port int) *Client {
	return &Client{
		client: http.DefaultClient,
		host:   fmt.Sprintf("%s:%d", host, port),
	}
}
