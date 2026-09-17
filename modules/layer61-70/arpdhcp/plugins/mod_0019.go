package arpdhcp

import (
    "time"
)

type arpdhcp0019 struct{}

func Newarpdhcp0019() *arpdhcp0019 {
    return &arpdhcp0019{}
}

func (e *arpdhcp0019) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0019) Name() string { return "arpdhcp0019" }
func (e *arpdhcp0019) Timestamp() time.Time { return time.Now() }
