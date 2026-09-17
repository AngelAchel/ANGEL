package arpdhcp

import (
    "time"
)

type arpdhcp0122 struct{}

func Newarpdhcp0122() *arpdhcp0122 {
    return &arpdhcp0122{}
}

func (e *arpdhcp0122) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0122) Name() string { return "arpdhcp0122" }
func (e *arpdhcp0122) Timestamp() time.Time { return time.Now() }
