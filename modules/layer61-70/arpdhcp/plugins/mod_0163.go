package arpdhcp

import (
    "time"
)

type arpdhcp0163 struct{}

func Newarpdhcp0163() *arpdhcp0163 {
    return &arpdhcp0163{}
}

func (e *arpdhcp0163) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0163) Name() string { return "arpdhcp0163" }
func (e *arpdhcp0163) Timestamp() time.Time { return time.Now() }
