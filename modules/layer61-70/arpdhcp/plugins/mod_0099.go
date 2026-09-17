package arpdhcp

import (
    "time"
)

type arpdhcp0099 struct{}

func Newarpdhcp0099() *arpdhcp0099 {
    return &arpdhcp0099{}
}

func (e *arpdhcp0099) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0099) Name() string { return "arpdhcp0099" }
func (e *arpdhcp0099) Timestamp() time.Time { return time.Now() }
