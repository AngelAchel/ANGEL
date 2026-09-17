package arpdhcp

import (
    "time"
)

type arpdhcp0073 struct{}

func Newarpdhcp0073() *arpdhcp0073 {
    return &arpdhcp0073{}
}

func (e *arpdhcp0073) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0073) Name() string { return "arpdhcp0073" }
func (e *arpdhcp0073) Timestamp() time.Time { return time.Now() }
