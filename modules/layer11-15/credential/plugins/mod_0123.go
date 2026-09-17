package credential

import (
	"time"
)

type credential0123 struct{}

func Newcredential0123() *credential0123 {
	return &credential0123{}
}

func (e *credential0123) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "credential:done")
	return results, nil
}

func (e *credential0123) Name() string { return "credential0123" }
func (e *credential0123) Timestamp() time.Time { return time.Now() }
