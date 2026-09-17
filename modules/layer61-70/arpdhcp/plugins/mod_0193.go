package arpdhcp

import (
    "time"
)

type arpdhcp0193 struct{}

func Newarpdhcp0193() *arpdhcp0193 {
    return &arpdhcp0193{}
}

func (e *arpdhcp0193) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0193) Name() string { return "arpdhcp0193" }
func (e *arpdhcp0193) Timestamp() time.Time { return time.Now() }
