package arpdhcp

import (
    "time"
)

type arpdhcp0164 struct{}

func Newarpdhcp0164() *arpdhcp0164 {
    return &arpdhcp0164{}
}

func (e *arpdhcp0164) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0164) Name() string { return "arpdhcp0164" }
func (e *arpdhcp0164) Timestamp() time.Time { return time.Now() }
