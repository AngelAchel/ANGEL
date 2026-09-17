package arpdhcp

import (
    "time"
)

type arpdhcp0147 struct{}

func Newarpdhcp0147() *arpdhcp0147 {
    return &arpdhcp0147{}
}

func (e *arpdhcp0147) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0147) Name() string { return "arpdhcp0147" }
func (e *arpdhcp0147) Timestamp() time.Time { return time.Now() }
