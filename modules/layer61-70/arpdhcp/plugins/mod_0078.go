package arpdhcp

import (
    "time"
)

type arpdhcp0078 struct{}

func Newarpdhcp0078() *arpdhcp0078 {
    return &arpdhcp0078{}
}

func (e *arpdhcp0078) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0078) Name() string { return "arpdhcp0078" }
func (e *arpdhcp0078) Timestamp() time.Time { return time.Now() }
