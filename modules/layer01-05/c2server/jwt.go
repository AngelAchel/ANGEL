package c2server

import (
	"time"
)

type Jwt struct{}

func NewJwt() *Jwt {
	return &Jwt{}
}

func (e *Jwt) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "jwt:done")
	return results, nil
}

func (e *Jwt) Name() string         { return "Jwt" }
func (e *Jwt) Timestamp() time.Time { return time.Now() }
