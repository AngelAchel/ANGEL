package arpdhcp

import (
    "time"
)

type arpdhcp0077 struct{}

func Newarpdhcp0077() *arpdhcp0077 {
    return &arpdhcp0077{}
}

func (e *arpdhcp0077) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0077) Name() string { return "arpdhcp0077" }
func (e *arpdhcp0077) Timestamp() time.Time { return time.Now() }
