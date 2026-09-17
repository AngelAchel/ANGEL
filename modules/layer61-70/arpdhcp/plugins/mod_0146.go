package arpdhcp

import (
    "time"
)

type arpdhcp0146 struct{}

func Newarpdhcp0146() *arpdhcp0146 {
    return &arpdhcp0146{}
}

func (e *arpdhcp0146) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0146) Name() string { return "arpdhcp0146" }
func (e *arpdhcp0146) Timestamp() time.Time { return time.Now() }
