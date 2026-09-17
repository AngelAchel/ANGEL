package c2server

import (
	"time"
)

type Nfs struct{}

func NewNfs() *Nfs {
	return &Nfs{}
}

func (e *Nfs) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "nfs:done")
	return results, nil
}

func (e *Nfs) Name() string { return "Nfs" }
func (e *Nfs) Timestamp() time.Time { return time.Now() }
