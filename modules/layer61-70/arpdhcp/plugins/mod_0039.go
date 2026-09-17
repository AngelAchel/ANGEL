package arpdhcp

import (
    "time"
)

type arpdhcp0039 struct{}

func Newarpdhcp0039() *arpdhcp0039 {
    return &arpdhcp0039{}
}

func (e *arpdhcp0039) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0039) Name() string { return "arpdhcp0039" }
func (e *arpdhcp0039) Timestamp() time.Time { return time.Now() }
