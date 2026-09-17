package arpdhcp

import (
    "time"
)

type arpdhcp0190 struct{}

func Newarpdhcp0190() *arpdhcp0190 {
    return &arpdhcp0190{}
}

func (e *arpdhcp0190) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0190) Name() string { return "arpdhcp0190" }
func (e *arpdhcp0190) Timestamp() time.Time { return time.Now() }
