package arpdhcp

import (
    "time"
)

type arpdhcp0023 struct{}

func Newarpdhcp0023() *arpdhcp0023 {
    return &arpdhcp0023{}
}

func (e *arpdhcp0023) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0023) Name() string { return "arpdhcp0023" }
func (e *arpdhcp0023) Timestamp() time.Time { return time.Now() }
