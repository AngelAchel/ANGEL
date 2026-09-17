package arpdhcp

import (
    "time"
)

type arpdhcp0142 struct{}

func Newarpdhcp0142() *arpdhcp0142 {
    return &arpdhcp0142{}
}

func (e *arpdhcp0142) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0142) Name() string { return "arpdhcp0142" }
func (e *arpdhcp0142) Timestamp() time.Time { return time.Now() }
