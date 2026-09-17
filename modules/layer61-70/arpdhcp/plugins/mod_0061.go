package arpdhcp

import (
    "time"
)

type arpdhcp0061 struct{}

func Newarpdhcp0061() *arpdhcp0061 {
    return &arpdhcp0061{}
}

func (e *arpdhcp0061) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0061) Name() string { return "arpdhcp0061" }
func (e *arpdhcp0061) Timestamp() time.Time { return time.Now() }
