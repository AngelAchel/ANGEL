package arpdhcp

import (
    "time"
)

type arpdhcp0010 struct{}

func Newarpdhcp0010() *arpdhcp0010 {
    return &arpdhcp0010{}
}

func (e *arpdhcp0010) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0010) Name() string { return "arpdhcp0010" }
func (e *arpdhcp0010) Timestamp() time.Time { return time.Now() }
