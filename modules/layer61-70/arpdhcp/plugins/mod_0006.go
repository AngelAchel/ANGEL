package arpdhcp

import (
    "time"
)

type arpdhcp0006 struct{}

func Newarpdhcp0006() *arpdhcp0006 {
    return &arpdhcp0006{}
}

func (e *arpdhcp0006) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0006) Name() string { return "arpdhcp0006" }
func (e *arpdhcp0006) Timestamp() time.Time { return time.Now() }
