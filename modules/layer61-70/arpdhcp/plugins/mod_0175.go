package arpdhcp

import (
    "time"
)

type arpdhcp0175 struct{}

func Newarpdhcp0175() *arpdhcp0175 {
    return &arpdhcp0175{}
}

func (e *arpdhcp0175) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0175) Name() string { return "arpdhcp0175" }
func (e *arpdhcp0175) Timestamp() time.Time { return time.Now() }
