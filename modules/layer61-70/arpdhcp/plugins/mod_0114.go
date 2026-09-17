package arpdhcp

import (
    "time"
)

type arpdhcp0114 struct{}

func Newarpdhcp0114() *arpdhcp0114 {
    return &arpdhcp0114{}
}

func (e *arpdhcp0114) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0114) Name() string { return "arpdhcp0114" }
func (e *arpdhcp0114) Timestamp() time.Time { return time.Now() }
