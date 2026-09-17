package arpdhcp

import (
    "time"
)

type arpdhcp0192 struct{}

func Newarpdhcp0192() *arpdhcp0192 {
    return &arpdhcp0192{}
}

func (e *arpdhcp0192) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0192) Name() string { return "arpdhcp0192" }
func (e *arpdhcp0192) Timestamp() time.Time { return time.Now() }
