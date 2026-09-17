package arpdhcp

import (
    "time"
)

type arpdhcp0041 struct{}

func Newarpdhcp0041() *arpdhcp0041 {
    return &arpdhcp0041{}
}

func (e *arpdhcp0041) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0041) Name() string { return "arpdhcp0041" }
func (e *arpdhcp0041) Timestamp() time.Time { return time.Now() }
