package arpdhcp

import (
    "time"
)

type arpdhcp0085 struct{}

func Newarpdhcp0085() *arpdhcp0085 {
    return &arpdhcp0085{}
}

func (e *arpdhcp0085) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0085) Name() string { return "arpdhcp0085" }
func (e *arpdhcp0085) Timestamp() time.Time { return time.Now() }
