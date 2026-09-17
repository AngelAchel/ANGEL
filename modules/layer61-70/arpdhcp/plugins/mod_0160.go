package arpdhcp

import (
    "time"
)

type arpdhcp0160 struct{}

func Newarpdhcp0160() *arpdhcp0160 {
    return &arpdhcp0160{}
}

func (e *arpdhcp0160) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0160) Name() string { return "arpdhcp0160" }
func (e *arpdhcp0160) Timestamp() time.Time { return time.Now() }
