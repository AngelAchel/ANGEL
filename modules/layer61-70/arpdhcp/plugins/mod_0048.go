package arpdhcp

import (
    "time"
)

type arpdhcp0048 struct{}

func Newarpdhcp0048() *arpdhcp0048 {
    return &arpdhcp0048{}
}

func (e *arpdhcp0048) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0048) Name() string { return "arpdhcp0048" }
func (e *arpdhcp0048) Timestamp() time.Time { return time.Now() }
