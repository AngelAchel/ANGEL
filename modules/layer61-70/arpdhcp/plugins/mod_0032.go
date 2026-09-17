package arpdhcp

import (
    "time"
)

type arpdhcp0032 struct{}

func Newarpdhcp0032() *arpdhcp0032 {
    return &arpdhcp0032{}
}

func (e *arpdhcp0032) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0032) Name() string { return "arpdhcp0032" }
func (e *arpdhcp0032) Timestamp() time.Time { return time.Now() }
