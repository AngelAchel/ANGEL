package arpdhcp

import (
    "time"
)

type arpdhcp0026 struct{}

func Newarpdhcp0026() *arpdhcp0026 {
    return &arpdhcp0026{}
}

func (e *arpdhcp0026) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0026) Name() string { return "arpdhcp0026" }
func (e *arpdhcp0026) Timestamp() time.Time { return time.Now() }
