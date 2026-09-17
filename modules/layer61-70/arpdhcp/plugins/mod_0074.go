package arpdhcp

import (
    "time"
)

type arpdhcp0074 struct{}

func Newarpdhcp0074() *arpdhcp0074 {
    return &arpdhcp0074{}
}

func (e *arpdhcp0074) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0074) Name() string { return "arpdhcp0074" }
func (e *arpdhcp0074) Timestamp() time.Time { return time.Now() }
