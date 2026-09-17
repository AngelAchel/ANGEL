package arpdhcp

import (
    "time"
)

type arpdhcp0120 struct{}

func Newarpdhcp0120() *arpdhcp0120 {
    return &arpdhcp0120{}
}

func (e *arpdhcp0120) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0120) Name() string { return "arpdhcp0120" }
func (e *arpdhcp0120) Timestamp() time.Time { return time.Now() }
