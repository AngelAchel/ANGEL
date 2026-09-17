package arpdhcp

import (
    "time"
)

type arpdhcp0156 struct{}

func Newarpdhcp0156() *arpdhcp0156 {
    return &arpdhcp0156{}
}

func (e *arpdhcp0156) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0156) Name() string { return "arpdhcp0156" }
func (e *arpdhcp0156) Timestamp() time.Time { return time.Now() }
