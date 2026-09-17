package storage

import (
	"time"
)

type Migrations struct{}

func NewMigrations() *Migrations {
	return &Migrations{}
}

func (m *Migrations) Migrate() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "migrations:done")
	return results, nil
}

func (m *Migrations) Name() string { return "Migrations" }
func (m *Migrations) Timestamp() time.Time { return time.Now() }
