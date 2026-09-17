package arpdhcp

import (
    "time"
)

type arpdhcp0145 struct{}

func Newarpdhcp0145() *arpdhcp0145 {
    return &arpdhcp0145{}
}

func (e *arpdhcp0145) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0145) Name() string { return "arpdhcp0145" }
func (e *arpdhcp0145) Timestamp() time.Time { return time.Now() }
