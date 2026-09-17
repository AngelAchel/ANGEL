package arpdhcp

import (
    "time"
)

type arpdhcp0025 struct{}

func Newarpdhcp0025() *arpdhcp0025 {
    return &arpdhcp0025{}
}

func (e *arpdhcp0025) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0025) Name() string { return "arpdhcp0025" }
func (e *arpdhcp0025) Timestamp() time.Time { return time.Now() }
