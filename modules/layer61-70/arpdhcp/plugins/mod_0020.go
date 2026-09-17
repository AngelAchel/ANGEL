package arpdhcp

import (
    "time"
)

type arpdhcp0020 struct{}

func Newarpdhcp0020() *arpdhcp0020 {
    return &arpdhcp0020{}
}

func (e *arpdhcp0020) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0020) Name() string { return "arpdhcp0020" }
func (e *arpdhcp0020) Timestamp() time.Time { return time.Now() }
