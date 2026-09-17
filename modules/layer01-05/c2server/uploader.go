package c2server

import (
	"time"
)

type Uploader struct{}

func NewUploader() *Uploader {
	return &Uploader{}
}

func (e *Uploader) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "uploader:done")
	return results, nil
}

func (e *Uploader) Name() string { return "Uploader" }
func (e *Uploader) Timestamp() time.Time { return time.Now() }
