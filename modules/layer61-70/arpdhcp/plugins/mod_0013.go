package arpdhcp

import (
    "time"
)

type arpdhcp0013 struct{}

func Newarpdhcp0013() *arpdhcp0013 {
    return &arpdhcp0013{}
}

func (e *arpdhcp0013) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0013) Name() string { return "arpdhcp0013" }
func (e *arpdhcp0013) Timestamp() time.Time { return time.Now() }
