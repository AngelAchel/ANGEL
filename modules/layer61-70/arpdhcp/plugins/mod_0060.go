package arpdhcp

import (
    "time"
)

type arpdhcp0060 struct{}

func Newarpdhcp0060() *arpdhcp0060 {
    return &arpdhcp0060{}
}

func (e *arpdhcp0060) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0060) Name() string { return "arpdhcp0060" }
func (e *arpdhcp0060) Timestamp() time.Time { return time.Now() }
