package arpdhcp

import (
    "time"
)

type arpdhcp0118 struct{}

func Newarpdhcp0118() *arpdhcp0118 {
    return &arpdhcp0118{}
}

func (e *arpdhcp0118) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0118) Name() string { return "arpdhcp0118" }
func (e *arpdhcp0118) Timestamp() time.Time { return time.Now() }
