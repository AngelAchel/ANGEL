package arpdhcp

import (
    "time"
)

type arpdhcp0015 struct{}

func Newarpdhcp0015() *arpdhcp0015 {
    return &arpdhcp0015{}
}

func (e *arpdhcp0015) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0015) Name() string { return "arpdhcp0015" }
func (e *arpdhcp0015) Timestamp() time.Time { return time.Now() }
