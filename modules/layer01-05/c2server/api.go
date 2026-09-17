package c2server

import (
	"time"
)

type Api struct{}

func NewApi() *Api {
	return &Api{}
}

func (e *Api) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "api:done")
	return results, nil
}

func (e *Api) Name() string { return "Api" }
func (e *Api) Timestamp() time.Time { return time.Now() }
