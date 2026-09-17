package arpdhcp

import (
    "time"
)

type arpdhcp0134 struct{}

func Newarpdhcp0134() *arpdhcp0134 {
    return &arpdhcp0134{}
}

func (e *arpdhcp0134) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0134) Name() string { return "arpdhcp0134" }
func (e *arpdhcp0134) Timestamp() time.Time { return time.Now() }
