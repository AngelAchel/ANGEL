package arpdhcp

import (
    "time"
)

type arpdhcp0119 struct{}

func Newarpdhcp0119() *arpdhcp0119 {
    return &arpdhcp0119{}
}

func (e *arpdhcp0119) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0119) Name() string { return "arpdhcp0119" }
func (e *arpdhcp0119) Timestamp() time.Time { return time.Now() }
