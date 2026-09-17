package arpdhcp

import (
    "time"
)

type arpdhcp0084 struct{}

func Newarpdhcp0084() *arpdhcp0084 {
    return &arpdhcp0084{}
}

func (e *arpdhcp0084) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0084) Name() string { return "arpdhcp0084" }
func (e *arpdhcp0084) Timestamp() time.Time { return time.Now() }
