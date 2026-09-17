package arpdhcp

import (
    "time"
)

type arpdhcp0185 struct{}

func Newarpdhcp0185() *arpdhcp0185 {
    return &arpdhcp0185{}
}

func (e *arpdhcp0185) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0185) Name() string { return "arpdhcp0185" }
func (e *arpdhcp0185) Timestamp() time.Time { return time.Now() }
