package arpdhcp

import (
    "time"
)

type arpdhcp0144 struct{}

func Newarpdhcp0144() *arpdhcp0144 {
    return &arpdhcp0144{}
}

func (e *arpdhcp0144) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0144) Name() string { return "arpdhcp0144" }
func (e *arpdhcp0144) Timestamp() time.Time { return time.Now() }
