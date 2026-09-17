package arpdhcp

import (
    "time"
)

type arpdhcp0101 struct{}

func Newarpdhcp0101() *arpdhcp0101 {
    return &arpdhcp0101{}
}

func (e *arpdhcp0101) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0101) Name() string { return "arpdhcp0101" }
func (e *arpdhcp0101) Timestamp() time.Time { return time.Now() }
