package c2server

import (
	"time"
)

type Flooder struct{}

func NewFlooder() *Flooder {
	return &Flooder{}
}

func (e *Flooder) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "flooder:done")
	return results, nil
}

func (e *Flooder) Name() string         { return "Flooder" }
func (e *Flooder) Timestamp() time.Time { return time.Now() }
