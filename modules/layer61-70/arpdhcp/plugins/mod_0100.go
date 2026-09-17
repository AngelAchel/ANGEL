package arpdhcp

import (
    "time"
)

type arpdhcp0100 struct{}

func Newarpdhcp0100() *arpdhcp0100 {
    return &arpdhcp0100{}
}

func (e *arpdhcp0100) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0100) Name() string { return "arpdhcp0100" }
func (e *arpdhcp0100) Timestamp() time.Time { return time.Now() }
