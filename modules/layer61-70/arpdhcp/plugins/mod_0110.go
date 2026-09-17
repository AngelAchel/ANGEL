package arpdhcp

import (
    "time"
)

type arpdhcp0110 struct{}

func Newarpdhcp0110() *arpdhcp0110 {
    return &arpdhcp0110{}
}

func (e *arpdhcp0110) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0110) Name() string { return "arpdhcp0110" }
func (e *arpdhcp0110) Timestamp() time.Time { return time.Now() }
