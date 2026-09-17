package arpdhcp

import (
    "time"
)

type arpdhcp0090 struct{}

func Newarpdhcp0090() *arpdhcp0090 {
    return &arpdhcp0090{}
}

func (e *arpdhcp0090) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0090) Name() string { return "arpdhcp0090" }
func (e *arpdhcp0090) Timestamp() time.Time { return time.Now() }
