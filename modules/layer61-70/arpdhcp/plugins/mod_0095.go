package arpdhcp

import (
    "time"
)

type arpdhcp0095 struct{}

func Newarpdhcp0095() *arpdhcp0095 {
    return &arpdhcp0095{}
}

func (e *arpdhcp0095) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0095) Name() string { return "arpdhcp0095" }
func (e *arpdhcp0095) Timestamp() time.Time { return time.Now() }
