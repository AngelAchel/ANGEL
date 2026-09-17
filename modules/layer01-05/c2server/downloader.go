package c2server

import (
	"time"
)

type Downloader struct{}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (e *Downloader) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "downloader:done")
	return results, nil
}

func (e *Downloader) Name() string         { return "Downloader" }
func (e *Downloader) Timestamp() time.Time { return time.Now() }
