package arpdhcp

import (
    "time"
)

type arpdhcp0103 struct{}

func Newarpdhcp0103() *arpdhcp0103 {
    return &arpdhcp0103{}
}

func (e *arpdhcp0103) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0103) Name() string { return "arpdhcp0103" }
func (e *arpdhcp0103) Timestamp() time.Time { return time.Now() }
