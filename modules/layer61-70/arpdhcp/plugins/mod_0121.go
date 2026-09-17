package arpdhcp

import (
    "time"
)

type arpdhcp0121 struct{}

func Newarpdhcp0121() *arpdhcp0121 {
    return &arpdhcp0121{}
}

func (e *arpdhcp0121) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0121) Name() string { return "arpdhcp0121" }
func (e *arpdhcp0121) Timestamp() time.Time { return time.Now() }
