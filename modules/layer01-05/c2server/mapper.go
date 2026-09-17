package c2server

import (
	"time"
)

type Mapper struct{}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (e *Mapper) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "mapper:done")
	return results, nil
}

func (e *Mapper) Name() string         { return "Mapper" }
func (e *Mapper) Timestamp() time.Time { return time.Now() }
