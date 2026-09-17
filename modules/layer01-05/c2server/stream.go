package c2server

import (
	"time"
)

type Stream struct{}

func NewStream() *Stream {
	return &Stream{}
}

func (e *Stream) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "stream:done")
	return results, nil
}

func (e *Stream) Name() string         { return "Stream" }
func (e *Stream) Timestamp() time.Time { return time.Now() }
