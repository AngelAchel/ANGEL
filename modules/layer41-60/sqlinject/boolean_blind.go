package sqlinject

import (
	"time"
)

type BooleanBlind struct{}

func NewBooleanBlind() *BooleanBlind {
	return &BooleanBlind{}
}

func (b *BooleanBlind) Inject() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "boolean_blind:done")
	return results, nil
}

func (b *BooleanBlind) Name() string         { return "BooleanBlind" }
func (b *BooleanBlind) Timestamp() time.Time { return time.Now() }
