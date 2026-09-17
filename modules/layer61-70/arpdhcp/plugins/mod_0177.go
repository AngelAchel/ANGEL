package arpdhcp

import (
    "time"
)

type arpdhcp0177 struct{}

func Newarpdhcp0177() *arpdhcp0177 {
    return &arpdhcp0177{}
}

func (e *arpdhcp0177) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0177) Name() string { return "arpdhcp0177" }
func (e *arpdhcp0177) Timestamp() time.Time { return time.Now() }
