package arpdhcp

import (
    "time"
)

type arpdhcp0138 struct{}

func Newarpdhcp0138() *arpdhcp0138 {
    return &arpdhcp0138{}
}

func (e *arpdhcp0138) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0138) Name() string { return "arpdhcp0138" }
func (e *arpdhcp0138) Timestamp() time.Time { return time.Now() }
