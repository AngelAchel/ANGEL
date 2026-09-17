package arpdhcp

import (
    "time"
)

type arpdhcp0088 struct{}

func Newarpdhcp0088() *arpdhcp0088 {
    return &arpdhcp0088{}
}

func (e *arpdhcp0088) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0088) Name() string { return "arpdhcp0088" }
func (e *arpdhcp0088) Timestamp() time.Time { return time.Now() }
