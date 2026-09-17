package arpdhcp

import (
    "time"
)

type arpdhcp0139 struct{}

func Newarpdhcp0139() *arpdhcp0139 {
    return &arpdhcp0139{}
}

func (e *arpdhcp0139) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0139) Name() string { return "arpdhcp0139" }
func (e *arpdhcp0139) Timestamp() time.Time { return time.Now() }
