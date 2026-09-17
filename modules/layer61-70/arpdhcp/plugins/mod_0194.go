package arpdhcp

import (
    "time"
)

type arpdhcp0194 struct{}

func Newarpdhcp0194() *arpdhcp0194 {
    return &arpdhcp0194{}
}

func (e *arpdhcp0194) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0194) Name() string { return "arpdhcp0194" }
func (e *arpdhcp0194) Timestamp() time.Time { return time.Now() }
