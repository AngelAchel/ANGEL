package arpdhcp

import (
    "time"
)

type arpdhcp0126 struct{}

func Newarpdhcp0126() *arpdhcp0126 {
    return &arpdhcp0126{}
}

func (e *arpdhcp0126) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0126) Name() string { return "arpdhcp0126" }
func (e *arpdhcp0126) Timestamp() time.Time { return time.Now() }
