package arpdhcp

import (
    "time"
)

type arpdhcp0116 struct{}

func Newarpdhcp0116() *arpdhcp0116 {
    return &arpdhcp0116{}
}

func (e *arpdhcp0116) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0116) Name() string { return "arpdhcp0116" }
func (e *arpdhcp0116) Timestamp() time.Time { return time.Now() }
