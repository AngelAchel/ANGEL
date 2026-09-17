package c2server

import (
	"time"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (e *Repository) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "repository:done")
	return results, nil
}

func (e *Repository) Name() string         { return "Repository" }
func (e *Repository) Timestamp() time.Time { return time.Now() }
