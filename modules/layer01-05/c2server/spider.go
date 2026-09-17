package c2server

import (
	"time"
)

type Spider struct{}

func NewSpider() *Spider {
	return &Spider{}
}

func (e *Spider) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "spider:done")
	return results, nil
}

func (e *Spider) Name() string { return "Spider" }
func (e *Spider) Timestamp() time.Time { return time.Now() }
