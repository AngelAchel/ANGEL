package arpdhcp

import (
    "time"
)

type arpdhcp0186 struct{}

func Newarpdhcp0186() *arpdhcp0186 {
    return &arpdhcp0186{}
}

func (e *arpdhcp0186) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0186) Name() string { return "arpdhcp0186" }
func (e *arpdhcp0186) Timestamp() time.Time { return time.Now() }
