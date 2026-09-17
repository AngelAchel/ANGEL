package arpdhcp

import (
    "time"
)

type arpdhcp0152 struct{}

func Newarpdhcp0152() *arpdhcp0152 {
    return &arpdhcp0152{}
}

func (e *arpdhcp0152) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0152) Name() string { return "arpdhcp0152" }
func (e *arpdhcp0152) Timestamp() time.Time { return time.Now() }
