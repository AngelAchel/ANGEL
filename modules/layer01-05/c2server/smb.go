package c2server

import (
	"time"
)

type Smb struct{}

func NewSmb() *Smb {
	return &Smb{}
}

func (e *Smb) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "smb:done")
	return results, nil
}

func (e *Smb) Name() string { return "Smb" }
func (e *Smb) Timestamp() time.Time { return time.Now() }
