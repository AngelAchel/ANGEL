package arpdhcp

import (
    "time"
)

type arpdhcp0087 struct{}

func Newarpdhcp0087() *arpdhcp0087 {
    return &arpdhcp0087{}
}

func (e *arpdhcp0087) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0087) Name() string { return "arpdhcp0087" }
func (e *arpdhcp0087) Timestamp() time.Time { return time.Now() }
