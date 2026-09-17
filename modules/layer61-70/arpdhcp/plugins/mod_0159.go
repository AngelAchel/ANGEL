package arpdhcp

import (
    "time"
)

type arpdhcp0159 struct{}

func Newarpdhcp0159() *arpdhcp0159 {
    return &arpdhcp0159{}
}

func (e *arpdhcp0159) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0159) Name() string { return "arpdhcp0159" }
func (e *arpdhcp0159) Timestamp() time.Time { return time.Now() }
