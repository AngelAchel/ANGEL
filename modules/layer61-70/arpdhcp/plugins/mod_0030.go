package arpdhcp

import (
    "time"
)

type arpdhcp0030 struct{}

func Newarpdhcp0030() *arpdhcp0030 {
    return &arpdhcp0030{}
}

func (e *arpdhcp0030) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0030) Name() string { return "arpdhcp0030" }
func (e *arpdhcp0030) Timestamp() time.Time { return time.Now() }
