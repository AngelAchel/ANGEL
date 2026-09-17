package arpdhcp

import (
    "time"
)

type arpdhcp0069 struct{}

func Newarpdhcp0069() *arpdhcp0069 {
    return &arpdhcp0069{}
}

func (e *arpdhcp0069) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0069) Name() string { return "arpdhcp0069" }
func (e *arpdhcp0069) Timestamp() time.Time { return time.Now() }
