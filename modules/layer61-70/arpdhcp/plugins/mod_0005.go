package arpdhcp

import (
    "time"
)

type arpdhcp0005 struct{}

func Newarpdhcp0005() *arpdhcp0005 {
    return &arpdhcp0005{}
}

func (e *arpdhcp0005) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0005) Name() string { return "arpdhcp0005" }
func (e *arpdhcp0005) Timestamp() time.Time { return time.Now() }
