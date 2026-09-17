package c2server

import (
	"time"
)

type Channel struct{}

func NewChannel() *Channel {
	return &Channel{}
}

func (e *Channel) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "channel:done")
	return results, nil
}

func (e *Channel) Name() string { return "Channel" }
func (e *Channel) Timestamp() time.Time { return time.Now() }
