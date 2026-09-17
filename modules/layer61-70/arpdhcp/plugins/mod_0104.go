package arpdhcp

import (
    "time"
)

type arpdhcp0104 struct{}

func Newarpdhcp0104() *arpdhcp0104 {
    return &arpdhcp0104{}
}

func (e *arpdhcp0104) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0104) Name() string { return "arpdhcp0104" }
func (e *arpdhcp0104) Timestamp() time.Time { return time.Now() }
