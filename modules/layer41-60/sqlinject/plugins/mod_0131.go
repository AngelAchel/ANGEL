package sqlinject

import (
	"time"
)

type Sqlinject0131 struct{}

func NewSqlinject0131() *Sqlinject0131 {
	return &Sqlinject0131{}
}

func (e *Sqlinject0131) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0131) Name() string { return "Sqlinject0131" }
func (e *Sqlinject0131) Timestamp() time.Time { return time.Now() }
