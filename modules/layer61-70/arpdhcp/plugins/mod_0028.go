package arpdhcp

import (
    "time"
)

type arpdhcp0028 struct{}

func Newarpdhcp0028() *arpdhcp0028 {
    return &arpdhcp0028{}
}

func (e *arpdhcp0028) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0028) Name() string { return "arpdhcp0028" }
func (e *arpdhcp0028) Timestamp() time.Time { return time.Now() }
