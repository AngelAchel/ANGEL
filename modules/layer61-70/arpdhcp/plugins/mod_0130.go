package arpdhcp

import (
    "time"
)

type arpdhcp0130 struct{}

func Newarpdhcp0130() *arpdhcp0130 {
    return &arpdhcp0130{}
}

func (e *arpdhcp0130) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0130) Name() string { return "arpdhcp0130" }
func (e *arpdhcp0130) Timestamp() time.Time { return time.Now() }
