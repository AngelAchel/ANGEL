package arpdhcp

import (
    "time"
)

type arpdhcp0008 struct{}

func Newarpdhcp0008() *arpdhcp0008 {
    return &arpdhcp0008{}
}

func (e *arpdhcp0008) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0008) Name() string { return "arpdhcp0008" }
func (e *arpdhcp0008) Timestamp() time.Time { return time.Now() }
