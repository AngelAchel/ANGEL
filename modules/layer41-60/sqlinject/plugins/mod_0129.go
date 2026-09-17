package sqlinject

import (
	"time"
)

type Sqlinject0129 struct{}

func NewSqlinject0129() *Sqlinject0129 {
	return &Sqlinject0129{}
}

func (e *Sqlinject0129) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0129) Name() string { return "Sqlinject0129" }
func (e *Sqlinject0129) Timestamp() time.Time { return time.Now() }
