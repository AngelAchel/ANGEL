package c2server

import (
	"time"
)

type Transformer struct{}

func NewTransformer() *Transformer {
	return &Transformer{}
}

func (e *Transformer) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "transformer:done")
	return results, nil
}

func (e *Transformer) Name() string         { return "Transformer" }
func (e *Transformer) Timestamp() time.Time { return time.Now() }
