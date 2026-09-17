package arpdhcp

import (
    "time"
)

type arpdhcp0096 struct{}

func Newarpdhcp0096() *arpdhcp0096 {
    return &arpdhcp0096{}
}

func (e *arpdhcp0096) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0096) Name() string { return "arpdhcp0096" }
func (e *arpdhcp0096) Timestamp() time.Time { return time.Now() }
