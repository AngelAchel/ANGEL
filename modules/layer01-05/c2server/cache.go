package c2server

import (
	"time"
)

type Cache struct{}

func NewCache() *Cache {
	return &Cache{}
}

func (e *Cache) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "cache:done")
	return results, nil
}

func (e *Cache) Name() string         { return "Cache" }
func (e *Cache) Timestamp() time.Time { return time.Now() }
