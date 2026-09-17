package arpdhcp

import (
    "time"
)

type arpdhcp0179 struct{}

func Newarpdhcp0179() *arpdhcp0179 {
    return &arpdhcp0179{}
}

func (e *arpdhcp0179) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0179) Name() string { return "arpdhcp0179" }
func (e *arpdhcp0179) Timestamp() time.Time { return time.Now() }
