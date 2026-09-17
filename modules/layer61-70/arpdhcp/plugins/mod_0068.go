package arpdhcp

import (
    "time"
)

type arpdhcp0068 struct{}

func Newarpdhcp0068() *arpdhcp0068 {
    return &arpdhcp0068{}
}

func (e *arpdhcp0068) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0068) Name() string { return "arpdhcp0068" }
func (e *arpdhcp0068) Timestamp() time.Time { return time.Now() }
