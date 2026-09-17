package arpdhcp

import (
    "time"
)

type arpdhcp0140 struct{}

func Newarpdhcp0140() *arpdhcp0140 {
    return &arpdhcp0140{}
}

func (e *arpdhcp0140) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0140) Name() string { return "arpdhcp0140" }
func (e *arpdhcp0140) Timestamp() time.Time { return time.Now() }
