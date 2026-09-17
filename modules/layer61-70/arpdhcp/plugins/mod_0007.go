package arpdhcp

import (
    "time"
)

type arpdhcp0007 struct{}

func Newarpdhcp0007() *arpdhcp0007 {
    return &arpdhcp0007{}
}

func (e *arpdhcp0007) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0007) Name() string { return "arpdhcp0007" }
func (e *arpdhcp0007) Timestamp() time.Time { return time.Now() }
