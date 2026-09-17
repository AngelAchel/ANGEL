package arpdhcp

import (
    "time"
)

type arpdhcp0071 struct{}

func Newarpdhcp0071() *arpdhcp0071 {
    return &arpdhcp0071{}
}

func (e *arpdhcp0071) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0071) Name() string { return "arpdhcp0071" }
func (e *arpdhcp0071) Timestamp() time.Time { return time.Now() }
