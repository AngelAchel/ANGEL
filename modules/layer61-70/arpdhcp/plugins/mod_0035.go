package arpdhcp

import (
    "time"
)

type arpdhcp0035 struct{}

func Newarpdhcp0035() *arpdhcp0035 {
    return &arpdhcp0035{}
}

func (e *arpdhcp0035) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0035) Name() string { return "arpdhcp0035" }
func (e *arpdhcp0035) Timestamp() time.Time { return time.Now() }
