package arpdhcp

import (
    "time"
)

type arpdhcp0133 struct{}

func Newarpdhcp0133() *arpdhcp0133 {
    return &arpdhcp0133{}
}

func (e *arpdhcp0133) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0133) Name() string { return "arpdhcp0133" }
func (e *arpdhcp0133) Timestamp() time.Time { return time.Now() }
