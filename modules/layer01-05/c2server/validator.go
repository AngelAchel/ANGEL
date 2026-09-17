package c2server

import (
	"time"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (e *Validator) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "validator:done")
	return results, nil
}

func (e *Validator) Name() string         { return "Validator" }
func (e *Validator) Timestamp() time.Time { return time.Now() }
