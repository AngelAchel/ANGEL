package arpdhcp

import (
    "time"
)

type arpdhcp0111 struct{}

func Newarpdhcp0111() *arpdhcp0111 {
    return &arpdhcp0111{}
}

func (e *arpdhcp0111) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0111) Name() string { return "arpdhcp0111" }
func (e *arpdhcp0111) Timestamp() time.Time { return time.Now() }
