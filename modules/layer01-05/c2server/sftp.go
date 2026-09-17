package c2server

import (
	"time"
)

type Sftp struct{}

func NewSftp() *Sftp {
	return &Sftp{}
}

func (e *Sftp) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sftp:done")
	return results, nil
}

func (e *Sftp) Name() string         { return "Sftp" }
func (e *Sftp) Timestamp() time.Time { return time.Now() }
