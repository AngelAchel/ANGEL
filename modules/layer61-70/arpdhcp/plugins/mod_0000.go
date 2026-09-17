package arpdhcp

import (
    "time"
)

type arpdhcp0000 struct{}

func Newarpdhcp0000() *arpdhcp0000 {
    return &arpdhcp0000{}
}

func (e *arpdhcp0000) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0000) Name() string { return "arpdhcp0000" }
func (e *arpdhcp0000) Timestamp() time.Time { return time.Now() }
