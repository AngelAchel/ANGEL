package arpdhcp

import (
    "time"
)

type arpdhcp0107 struct{}

func Newarpdhcp0107() *arpdhcp0107 {
    return &arpdhcp0107{}
}

func (e *arpdhcp0107) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0107) Name() string { return "arpdhcp0107" }
func (e *arpdhcp0107) Timestamp() time.Time { return time.Now() }
