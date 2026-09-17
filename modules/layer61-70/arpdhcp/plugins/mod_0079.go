package arpdhcp

import (
    "time"
)

type arpdhcp0079 struct{}

func Newarpdhcp0079() *arpdhcp0079 {
    return &arpdhcp0079{}
}

func (e *arpdhcp0079) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0079) Name() string { return "arpdhcp0079" }
func (e *arpdhcp0079) Timestamp() time.Time { return time.Now() }
