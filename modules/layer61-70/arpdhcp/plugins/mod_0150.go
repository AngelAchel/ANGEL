package arpdhcp

import (
    "time"
)

type arpdhcp0150 struct{}

func Newarpdhcp0150() *arpdhcp0150 {
    return &arpdhcp0150{}
}

func (e *arpdhcp0150) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0150) Name() string { return "arpdhcp0150" }
func (e *arpdhcp0150) Timestamp() time.Time { return time.Now() }
