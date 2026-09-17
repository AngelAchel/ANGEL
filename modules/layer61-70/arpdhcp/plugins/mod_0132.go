package arpdhcp

import (
    "time"
)

type arpdhcp0132 struct{}

func Newarpdhcp0132() *arpdhcp0132 {
    return &arpdhcp0132{}
}

func (e *arpdhcp0132) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0132) Name() string { return "arpdhcp0132" }
func (e *arpdhcp0132) Timestamp() time.Time { return time.Now() }
