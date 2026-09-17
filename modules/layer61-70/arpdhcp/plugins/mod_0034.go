package arpdhcp

import (
    "time"
)

type arpdhcp0034 struct{}

func Newarpdhcp0034() *arpdhcp0034 {
    return &arpdhcp0034{}
}

func (e *arpdhcp0034) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0034) Name() string { return "arpdhcp0034" }
func (e *arpdhcp0034) Timestamp() time.Time { return time.Now() }
