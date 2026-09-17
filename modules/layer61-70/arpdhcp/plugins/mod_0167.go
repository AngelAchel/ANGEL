package arpdhcp

import (
    "time"
)

type arpdhcp0167 struct{}

func Newarpdhcp0167() *arpdhcp0167 {
    return &arpdhcp0167{}
}

func (e *arpdhcp0167) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0167) Name() string { return "arpdhcp0167" }
func (e *arpdhcp0167) Timestamp() time.Time { return time.Now() }
