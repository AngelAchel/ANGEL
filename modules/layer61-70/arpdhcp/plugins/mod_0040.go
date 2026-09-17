package arpdhcp

import (
    "time"
)

type arpdhcp0040 struct{}

func Newarpdhcp0040() *arpdhcp0040 {
    return &arpdhcp0040{}
}

func (e *arpdhcp0040) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0040) Name() string { return "arpdhcp0040" }
func (e *arpdhcp0040) Timestamp() time.Time { return time.Now() }
