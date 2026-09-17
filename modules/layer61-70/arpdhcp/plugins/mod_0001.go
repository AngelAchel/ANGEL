package arpdhcp

import (
    "time"
)

type arpdhcp0001 struct{}

func Newarpdhcp0001() *arpdhcp0001 {
    return &arpdhcp0001{}
}

func (e *arpdhcp0001) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0001) Name() string { return "arpdhcp0001" }
func (e *arpdhcp0001) Timestamp() time.Time { return time.Now() }
