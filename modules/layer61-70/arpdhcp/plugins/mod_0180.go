package arpdhcp

import (
    "time"
)

type arpdhcp0180 struct{}

func Newarpdhcp0180() *arpdhcp0180 {
    return &arpdhcp0180{}
}

func (e *arpdhcp0180) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0180) Name() string { return "arpdhcp0180" }
func (e *arpdhcp0180) Timestamp() time.Time { return time.Now() }
