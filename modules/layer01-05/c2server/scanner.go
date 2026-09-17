package c2server

import (
	"time"
)

type Scanner struct{}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (e *Scanner) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "scanner:done")
	return results, nil
}

func (e *Scanner) Name() string         { return "Scanner" }
func (e *Scanner) Timestamp() time.Time { return time.Now() }
