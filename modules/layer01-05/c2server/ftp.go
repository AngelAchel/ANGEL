package c2server

import (
	"time"
)

type Ftp struct{}

func NewFtp() *Ftp {
	return &Ftp{}
}

func (e *Ftp) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "ftp:done")
	return results, nil
}

func (e *Ftp) Name() string { return "Ftp" }
func (e *Ftp) Timestamp() time.Time { return time.Now() }
