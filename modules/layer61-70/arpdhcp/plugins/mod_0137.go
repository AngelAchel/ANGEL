package arpdhcp

import (
    "time"
)

type arpdhcp0137 struct{}

func Newarpdhcp0137() *arpdhcp0137 {
    return &arpdhcp0137{}
}

func (e *arpdhcp0137) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0137) Name() string { return "arpdhcp0137" }
func (e *arpdhcp0137) Timestamp() time.Time { return time.Now() }
