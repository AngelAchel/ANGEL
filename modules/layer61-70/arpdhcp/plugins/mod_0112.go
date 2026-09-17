package arpdhcp

import (
    "time"
)

type arpdhcp0112 struct{}

func Newarpdhcp0112() *arpdhcp0112 {
    return &arpdhcp0112{}
}

func (e *arpdhcp0112) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0112) Name() string { return "arpdhcp0112" }
func (e *arpdhcp0112) Timestamp() time.Time { return time.Now() }
