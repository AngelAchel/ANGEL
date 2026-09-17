package arpdhcp

import (
    "time"
)

type arpdhcp0157 struct{}

func Newarpdhcp0157() *arpdhcp0157 {
    return &arpdhcp0157{}
}

func (e *arpdhcp0157) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0157) Name() string { return "arpdhcp0157" }
func (e *arpdhcp0157) Timestamp() time.Time { return time.Now() }
