package evasion

import (
	"time"
)

type BooleanBlind struct{}

func NewBooleanBlind() *BooleanBlind {
	return &BooleanBlind{}
}

func (e *BooleanBlind) Inject(query string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "boolean_blind:injected")
	return results, nil
}

func (e *BooleanBlind) Name() string { return "BooleanBlind" }
func (e *BooleanBlind) Category() EvasionCategory { return CategoryInjection }
func (e *BooleanBlind) Timestamp() time.Time { return time.Now() }
