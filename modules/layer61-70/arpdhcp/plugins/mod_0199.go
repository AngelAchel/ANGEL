package arpdhcp

import (
    "time"
)

type arpdhcp0199 struct{}

func Newarpdhcp0199() *arpdhcp0199 {
    return &arpdhcp0199{}
}

func (e *arpdhcp0199) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0199) Name() string { return "arpdhcp0199" }
func (e *arpdhcp0199) Timestamp() time.Time { return time.Now() }
