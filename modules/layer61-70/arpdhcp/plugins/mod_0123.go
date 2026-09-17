package arpdhcp

import (
    "time"
)

type arpdhcp0123 struct{}

func Newarpdhcp0123() *arpdhcp0123 {
    return &arpdhcp0123{}
}

func (e *arpdhcp0123) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0123) Name() string { return "arpdhcp0123" }
func (e *arpdhcp0123) Timestamp() time.Time { return time.Now() }
