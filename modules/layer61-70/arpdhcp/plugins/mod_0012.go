package arpdhcp

import (
    "time"
)

type arpdhcp0012 struct{}

func Newarpdhcp0012() *arpdhcp0012 {
    return &arpdhcp0012{}
}

func (e *arpdhcp0012) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0012) Name() string { return "arpdhcp0012" }
func (e *arpdhcp0012) Timestamp() time.Time { return time.Now() }
