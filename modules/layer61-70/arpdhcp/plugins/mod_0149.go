package arpdhcp

import (
    "time"
)

type arpdhcp0149 struct{}

func Newarpdhcp0149() *arpdhcp0149 {
    return &arpdhcp0149{}
}

func (e *arpdhcp0149) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0149) Name() string { return "arpdhcp0149" }
func (e *arpdhcp0149) Timestamp() time.Time { return time.Now() }
