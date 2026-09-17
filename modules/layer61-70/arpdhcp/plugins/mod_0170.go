package arpdhcp

import (
    "time"
)

type arpdhcp0170 struct{}

func Newarpdhcp0170() *arpdhcp0170 {
    return &arpdhcp0170{}
}

func (e *arpdhcp0170) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0170) Name() string { return "arpdhcp0170" }
func (e *arpdhcp0170) Timestamp() time.Time { return time.Now() }
