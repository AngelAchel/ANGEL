package arpdhcp

import (
    "time"
)

type arpdhcp0102 struct{}

func Newarpdhcp0102() *arpdhcp0102 {
    return &arpdhcp0102{}
}

func (e *arpdhcp0102) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0102) Name() string { return "arpdhcp0102" }
func (e *arpdhcp0102) Timestamp() time.Time { return time.Now() }
