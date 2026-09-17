package c2server

import (
	"time"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (e *Client) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "client:done")
	return results, nil
}

func (e *Client) Name() string { return "Client" }
func (e *Client) Timestamp() time.Time { return time.Now() }
