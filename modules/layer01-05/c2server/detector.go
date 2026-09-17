package c2server

import (
	"time"
)

type Detector struct{}

func NewDetector() *Detector {
	return &Detector{}
}

func (e *Detector) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "detector:done")
	return results, nil
}

func (e *Detector) Name() string         { return "Detector" }
func (e *Detector) Timestamp() time.Time { return time.Now() }
