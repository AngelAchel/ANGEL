package arpdhcp

import (
    "time"
)

type arpdhcp0125 struct{}

func Newarpdhcp0125() *arpdhcp0125 {
    return &arpdhcp0125{}
}

func (e *arpdhcp0125) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0125) Name() string { return "arpdhcp0125" }
func (e *arpdhcp0125) Timestamp() time.Time { return time.Now() }
