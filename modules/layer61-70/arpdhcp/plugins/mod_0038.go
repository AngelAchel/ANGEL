package arpdhcp

import (
    "time"
)

type arpdhcp0038 struct{}

func Newarpdhcp0038() *arpdhcp0038 {
    return &arpdhcp0038{}
}

func (e *arpdhcp0038) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0038) Name() string { return "arpdhcp0038" }
func (e *arpdhcp0038) Timestamp() time.Time { return time.Now() }
