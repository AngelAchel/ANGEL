package arpdhcp

import (
    "time"
)

type arpdhcp0108 struct{}

func Newarpdhcp0108() *arpdhcp0108 {
    return &arpdhcp0108{}
}

func (e *arpdhcp0108) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0108) Name() string { return "arpdhcp0108" }
func (e *arpdhcp0108) Timestamp() time.Time { return time.Now() }
