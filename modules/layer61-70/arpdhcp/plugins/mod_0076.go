package arpdhcp

import (
    "time"
)

type arpdhcp0076 struct{}

func Newarpdhcp0076() *arpdhcp0076 {
    return &arpdhcp0076{}
}

func (e *arpdhcp0076) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0076) Name() string { return "arpdhcp0076" }
func (e *arpdhcp0076) Timestamp() time.Time { return time.Now() }
