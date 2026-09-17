package arpdhcp

import (
    "time"
)

type arpdhcp0017 struct{}

func Newarpdhcp0017() *arpdhcp0017 {
    return &arpdhcp0017{}
}

func (e *arpdhcp0017) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0017) Name() string { return "arpdhcp0017" }
func (e *arpdhcp0017) Timestamp() time.Time { return time.Now() }
