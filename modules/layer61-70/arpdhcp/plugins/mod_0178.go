package arpdhcp

import (
    "time"
)

type arpdhcp0178 struct{}

func Newarpdhcp0178() *arpdhcp0178 {
    return &arpdhcp0178{}
}

func (e *arpdhcp0178) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0178) Name() string { return "arpdhcp0178" }
func (e *arpdhcp0178) Timestamp() time.Time { return time.Now() }
