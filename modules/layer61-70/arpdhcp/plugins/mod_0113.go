package arpdhcp

import (
    "time"
)

type arpdhcp0113 struct{}

func Newarpdhcp0113() *arpdhcp0113 {
    return &arpdhcp0113{}
}

func (e *arpdhcp0113) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0113) Name() string { return "arpdhcp0113" }
func (e *arpdhcp0113) Timestamp() time.Time { return time.Now() }
