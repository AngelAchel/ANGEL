package arpdhcp

import (
    "time"
)

type arpdhcp0089 struct{}

func Newarpdhcp0089() *arpdhcp0089 {
    return &arpdhcp0089{}
}

func (e *arpdhcp0089) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0089) Name() string { return "arpdhcp0089" }
func (e *arpdhcp0089) Timestamp() time.Time { return time.Now() }
