package evasion

import (
	"time"
)

type DoubleURL struct{}

func NewDoubleURL() *DoubleURL {
	return &DoubleURL{}
}

func (e *DoubleURL) Encode(url string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "double_url:encoded")
	return results, nil
}

func (e *DoubleURL) Name() string { return "DoubleURL" }
func (e *DoubleURL) Timestamp() time.Time { return time.Now() }
