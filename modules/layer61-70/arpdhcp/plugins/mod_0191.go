package arpdhcp

import (
    "time"
)

type arpdhcp0191 struct{}

func Newarpdhcp0191() *arpdhcp0191 {
    return &arpdhcp0191{}
}

func (e *arpdhcp0191) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0191) Name() string { return "arpdhcp0191" }
func (e *arpdhcp0191) Timestamp() time.Time { return time.Now() }
