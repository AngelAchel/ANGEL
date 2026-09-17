package arpdhcp

import (
    "time"
)

type arpdhcp0045 struct{}

func Newarpdhcp0045() *arpdhcp0045 {
    return &arpdhcp0045{}
}

func (e *arpdhcp0045) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0045) Name() string { return "arpdhcp0045" }
func (e *arpdhcp0045) Timestamp() time.Time { return time.Now() }
