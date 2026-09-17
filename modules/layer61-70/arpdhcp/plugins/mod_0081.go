package arpdhcp

import (
    "time"
)

type arpdhcp0081 struct{}

func Newarpdhcp0081() *arpdhcp0081 {
    return &arpdhcp0081{}
}

func (e *arpdhcp0081) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0081) Name() string { return "arpdhcp0081" }
func (e *arpdhcp0081) Timestamp() time.Time { return time.Now() }
