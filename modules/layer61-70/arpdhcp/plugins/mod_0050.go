package arpdhcp

import (
    "time"
)

type arpdhcp0050 struct{}

func Newarpdhcp0050() *arpdhcp0050 {
    return &arpdhcp0050{}
}

func (e *arpdhcp0050) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0050) Name() string { return "arpdhcp0050" }
func (e *arpdhcp0050) Timestamp() time.Time { return time.Now() }
