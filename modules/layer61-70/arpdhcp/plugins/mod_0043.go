package arpdhcp

import (
    "time"
)

type arpdhcp0043 struct{}

func Newarpdhcp0043() *arpdhcp0043 {
    return &arpdhcp0043{}
}

func (e *arpdhcp0043) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0043) Name() string { return "arpdhcp0043" }
func (e *arpdhcp0043) Timestamp() time.Time { return time.Now() }
