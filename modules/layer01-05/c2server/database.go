package c2server

import (
	"time"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (e *Database) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "database:done")
	return results, nil
}

func (e *Database) Name() string { return "Database" }
func (e *Database) Timestamp() time.Time { return time.Now() }
