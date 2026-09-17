package arpdhcp

import (
    "time"
)

type arpdhcp0082 struct{}

func Newarpdhcp0082() *arpdhcp0082 {
    return &arpdhcp0082{}
}

func (e *arpdhcp0082) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0082) Name() string { return "arpdhcp0082" }
func (e *arpdhcp0082) Timestamp() time.Time { return time.Now() }
