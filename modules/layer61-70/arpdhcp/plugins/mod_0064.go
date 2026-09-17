package arpdhcp

import (
    "time"
)

type arpdhcp0064 struct{}

func Newarpdhcp0064() *arpdhcp0064 {
    return &arpdhcp0064{}
}

func (e *arpdhcp0064) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0064) Name() string { return "arpdhcp0064" }
func (e *arpdhcp0064) Timestamp() time.Time { return time.Now() }
