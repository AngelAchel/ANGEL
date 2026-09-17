package arpdhcp

import (
    "time"
)

type arpdhcp0188 struct{}

func Newarpdhcp0188() *arpdhcp0188 {
    return &arpdhcp0188{}
}

func (e *arpdhcp0188) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0188) Name() string { return "arpdhcp0188" }
func (e *arpdhcp0188) Timestamp() time.Time { return time.Now() }
