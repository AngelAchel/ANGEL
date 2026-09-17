package arpdhcp

import (
    "time"
)

type arpdhcp0094 struct{}

func Newarpdhcp0094() *arpdhcp0094 {
    return &arpdhcp0094{}
}

func (e *arpdhcp0094) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0094) Name() string { return "arpdhcp0094" }
func (e *arpdhcp0094) Timestamp() time.Time { return time.Now() }
