package arpdhcp

import (
    "time"
)

type arpdhcp0003 struct{}

func Newarpdhcp0003() *arpdhcp0003 {
    return &arpdhcp0003{}
}

func (e *arpdhcp0003) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0003) Name() string { return "arpdhcp0003" }
func (e *arpdhcp0003) Timestamp() time.Time { return time.Now() }
