package arpdhcp

import (
    "time"
)

type arpdhcp0169 struct{}

func Newarpdhcp0169() *arpdhcp0169 {
    return &arpdhcp0169{}
}

func (e *arpdhcp0169) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0169) Name() string { return "arpdhcp0169" }
func (e *arpdhcp0169) Timestamp() time.Time { return time.Now() }
