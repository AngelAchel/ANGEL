package arpdhcp

import (
    "time"
)

type arpdhcp0031 struct{}

func Newarpdhcp0031() *arpdhcp0031 {
    return &arpdhcp0031{}
}

func (e *arpdhcp0031) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0031) Name() string { return "arpdhcp0031" }
func (e *arpdhcp0031) Timestamp() time.Time { return time.Now() }
