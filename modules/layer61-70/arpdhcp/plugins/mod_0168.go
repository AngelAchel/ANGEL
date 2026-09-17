package arpdhcp

import (
    "time"
)

type arpdhcp0168 struct{}

func Newarpdhcp0168() *arpdhcp0168 {
    return &arpdhcp0168{}
}

func (e *arpdhcp0168) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0168) Name() string { return "arpdhcp0168" }
func (e *arpdhcp0168) Timestamp() time.Time { return time.Now() }
