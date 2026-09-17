package arpdhcp

import (
    "time"
)

type arpdhcp0083 struct{}

func Newarpdhcp0083() *arpdhcp0083 {
    return &arpdhcp0083{}
}

func (e *arpdhcp0083) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0083) Name() string { return "arpdhcp0083" }
func (e *arpdhcp0083) Timestamp() time.Time { return time.Now() }
