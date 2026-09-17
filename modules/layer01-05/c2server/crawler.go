package c2server

import (
	"time"
)

type Crawler struct{}

func NewCrawler() *Crawler {
	return &Crawler{}
}

func (e *Crawler) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "crawler:done")
	return results, nil
}

func (e *Crawler) Name() string { return "Crawler" }
func (e *Crawler) Timestamp() time.Time { return time.Now() }
