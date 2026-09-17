package c2server

import (
	"time"
)

type Persister struct{}

func NewPersister() *Persister {
	return &Persister{}
}

func (e *Persister) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "persister:done")
	return results, nil
}

func (e *Persister) Name() string         { return "Persister" }
func (e *Persister) Timestamp() time.Time { return time.Now() }
