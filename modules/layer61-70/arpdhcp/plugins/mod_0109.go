package arpdhcp

import (
    "time"
)

type arpdhcp0109 struct{}

func Newarpdhcp0109() *arpdhcp0109 {
    return &arpdhcp0109{}
}

func (e *arpdhcp0109) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0109) Name() string { return "arpdhcp0109" }
func (e *arpdhcp0109) Timestamp() time.Time { return time.Now() }
