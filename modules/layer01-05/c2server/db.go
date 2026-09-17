package c2server

import (
	"time"
)

type Db struct{}

func NewDb() *Db {
	return &Db{}
}

func (e *Db) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "db:done")
	return results, nil
}

func (e *Db) Name() string         { return "Db" }
func (e *Db) Timestamp() time.Time { return time.Now() }
