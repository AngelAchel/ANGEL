package arpdhcp

import (
    "time"
)

type arpdhcp0198 struct{}

func Newarpdhcp0198() *arpdhcp0198 {
    return &arpdhcp0198{}
}

func (e *arpdhcp0198) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0198) Name() string { return "arpdhcp0198" }
func (e *arpdhcp0198) Timestamp() time.Time { return time.Now() }
