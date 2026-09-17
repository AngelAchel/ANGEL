package arpdhcp

import (
    "time"
)

type arpdhcp0033 struct{}

func Newarpdhcp0033() *arpdhcp0033 {
    return &arpdhcp0033{}
}

func (e *arpdhcp0033) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0033) Name() string { return "arpdhcp0033" }
func (e *arpdhcp0033) Timestamp() time.Time { return time.Now() }
