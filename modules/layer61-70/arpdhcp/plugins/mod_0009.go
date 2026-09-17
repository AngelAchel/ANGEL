package arpdhcp

import (
    "time"
)

type arpdhcp0009 struct{}

func Newarpdhcp0009() *arpdhcp0009 {
    return &arpdhcp0009{}
}

func (e *arpdhcp0009) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0009) Name() string { return "arpdhcp0009" }
func (e *arpdhcp0009) Timestamp() time.Time { return time.Now() }
