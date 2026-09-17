package arpdhcp

import (
    "time"
)

type arpdhcp0044 struct{}

func Newarpdhcp0044() *arpdhcp0044 {
    return &arpdhcp0044{}
}

func (e *arpdhcp0044) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0044) Name() string { return "arpdhcp0044" }
func (e *arpdhcp0044) Timestamp() time.Time { return time.Now() }
