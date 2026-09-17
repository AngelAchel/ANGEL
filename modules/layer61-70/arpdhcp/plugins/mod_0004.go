package arpdhcp

import (
    "time"
)

type arpdhcp0004 struct{}

func Newarpdhcp0004() *arpdhcp0004 {
    return &arpdhcp0004{}
}

func (e *arpdhcp0004) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0004) Name() string { return "arpdhcp0004" }
func (e *arpdhcp0004) Timestamp() time.Time { return time.Now() }
