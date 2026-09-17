package c2server

import (
	"time"
)

type Store struct{}

func NewStore() *Store {
	return &Store{}
}

func (e *Store) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "store:done")
	return results, nil
}

func (e *Store) Name() string { return "Store" }
func (e *Store) Timestamp() time.Time { return time.Now() }
