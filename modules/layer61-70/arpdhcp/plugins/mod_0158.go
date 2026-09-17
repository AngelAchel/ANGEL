package arpdhcp

import (
    "time"
)

type arpdhcp0158 struct{}

func Newarpdhcp0158() *arpdhcp0158 {
    return &arpdhcp0158{}
}

func (e *arpdhcp0158) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0158) Name() string { return "arpdhcp0158" }
func (e *arpdhcp0158) Timestamp() time.Time { return time.Now() }
