package arpdhcp

import (
    "time"
)

type arpdhcp0037 struct{}

func Newarpdhcp0037() *arpdhcp0037 {
    return &arpdhcp0037{}
}

func (e *arpdhcp0037) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0037) Name() string { return "arpdhcp0037" }
func (e *arpdhcp0037) Timestamp() time.Time { return time.Now() }
