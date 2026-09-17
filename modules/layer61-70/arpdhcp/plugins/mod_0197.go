package arpdhcp

import (
    "time"
)

type arpdhcp0197 struct{}

func Newarpdhcp0197() *arpdhcp0197 {
    return &arpdhcp0197{}
}

func (e *arpdhcp0197) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0197) Name() string { return "arpdhcp0197" }
func (e *arpdhcp0197) Timestamp() time.Time { return time.Now() }
