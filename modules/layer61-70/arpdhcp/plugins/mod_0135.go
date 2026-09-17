package arpdhcp

import (
    "time"
)

type arpdhcp0135 struct{}

func Newarpdhcp0135() *arpdhcp0135 {
    return &arpdhcp0135{}
}

func (e *arpdhcp0135) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0135) Name() string { return "arpdhcp0135" }
func (e *arpdhcp0135) Timestamp() time.Time { return time.Now() }
