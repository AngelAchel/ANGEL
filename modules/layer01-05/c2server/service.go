package c2server

import (
	"time"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (e *Service) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "service:done")
	return results, nil
}

func (e *Service) Name() string         { return "Service" }
func (e *Service) Timestamp() time.Time { return time.Now() }
