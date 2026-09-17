package c2server

import (
	"time"
)

type Oauth struct{}

func NewOauth() *Oauth {
	return &Oauth{}
}

func (e *Oauth) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "oauth:done")
	return results, nil
}

func (e *Oauth) Name() string         { return "Oauth" }
func (e *Oauth) Timestamp() time.Time { return time.Now() }
