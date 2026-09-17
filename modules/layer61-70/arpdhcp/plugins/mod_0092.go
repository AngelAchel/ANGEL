package arpdhcp

import (
    "time"
)

type arpdhcp0092 struct{}

func Newarpdhcp0092() *arpdhcp0092 {
    return &arpdhcp0092{}
}

func (e *arpdhcp0092) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0092) Name() string { return "arpdhcp0092" }
func (e *arpdhcp0092) Timestamp() time.Time { return time.Now() }
