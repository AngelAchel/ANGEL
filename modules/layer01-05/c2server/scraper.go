package c2server

import (
	"time"
)

type Scraper struct{}

func NewScraper() *Scraper {
	return &Scraper{}
}

func (e *Scraper) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "scraper:done")
	return results, nil
}

func (e *Scraper) Name() string         { return "Scraper" }
func (e *Scraper) Timestamp() time.Time { return time.Now() }
