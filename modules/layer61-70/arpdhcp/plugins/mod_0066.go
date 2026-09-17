package arpdhcp

import (
    "time"
)

type arpdhcp0066 struct{}

func Newarpdhcp0066() *arpdhcp0066 {
    return &arpdhcp0066{}
}

func (e *arpdhcp0066) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0066) Name() string { return "arpdhcp0066" }
func (e *arpdhcp0066) Timestamp() time.Time { return time.Now() }
