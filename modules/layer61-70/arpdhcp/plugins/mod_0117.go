package arpdhcp

import (
    "time"
)

type arpdhcp0117 struct{}

func Newarpdhcp0117() *arpdhcp0117 {
    return &arpdhcp0117{}
}

func (e *arpdhcp0117) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0117) Name() string { return "arpdhcp0117" }
func (e *arpdhcp0117) Timestamp() time.Time { return time.Now() }
