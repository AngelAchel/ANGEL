package arpdhcp

import (
    "time"
)

type arpdhcp0016 struct{}

func Newarpdhcp0016() *arpdhcp0016 {
    return &arpdhcp0016{}
}

func (e *arpdhcp0016) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0016) Name() string { return "arpdhcp0016" }
func (e *arpdhcp0016) Timestamp() time.Time { return time.Now() }
