package arpdhcp

import (
    "time"
)

type arpdhcp0018 struct{}

func Newarpdhcp0018() *arpdhcp0018 {
    return &arpdhcp0018{}
}

func (e *arpdhcp0018) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0018) Name() string { return "arpdhcp0018" }
func (e *arpdhcp0018) Timestamp() time.Time { return time.Now() }
