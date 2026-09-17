package arpdhcp

import (
    "time"
)

type arpdhcp0022 struct{}

func Newarpdhcp0022() *arpdhcp0022 {
    return &arpdhcp0022{}
}

func (e *arpdhcp0022) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0022) Name() string { return "arpdhcp0022" }
func (e *arpdhcp0022) Timestamp() time.Time { return time.Now() }
