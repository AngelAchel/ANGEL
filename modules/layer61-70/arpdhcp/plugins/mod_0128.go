package arpdhcp

import (
    "time"
)

type arpdhcp0128 struct{}

func Newarpdhcp0128() *arpdhcp0128 {
    return &arpdhcp0128{}
}

func (e *arpdhcp0128) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0128) Name() string { return "arpdhcp0128" }
func (e *arpdhcp0128) Timestamp() time.Time { return time.Now() }
