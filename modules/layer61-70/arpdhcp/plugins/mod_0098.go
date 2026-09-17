package arpdhcp

import (
    "time"
)

type arpdhcp0098 struct{}

func Newarpdhcp0098() *arpdhcp0098 {
    return &arpdhcp0098{}
}

func (e *arpdhcp0098) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0098) Name() string { return "arpdhcp0098" }
func (e *arpdhcp0098) Timestamp() time.Time { return time.Now() }
