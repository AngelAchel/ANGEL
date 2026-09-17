package c2

import (
	"time"
)

type c20120 struct{}

func Newc20120() *c20120 {
	return &c20120{}
}

func (e *c20120) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2:done")
	return results, nil
}

func (e *c20120) Name() string { return "c20120" }
func (e *c20120) Timestamp() time.Time { return time.Now() }
