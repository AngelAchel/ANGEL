package arpdhcp

import (
    "time"
)

type arpdhcp0070 struct{}

func Newarpdhcp0070() *arpdhcp0070 {
    return &arpdhcp0070{}
}

func (e *arpdhcp0070) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0070) Name() string { return "arpdhcp0070" }
func (e *arpdhcp0070) Timestamp() time.Time { return time.Now() }
