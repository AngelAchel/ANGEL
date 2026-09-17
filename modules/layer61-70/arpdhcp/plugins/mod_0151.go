package arpdhcp

import (
    "time"
)

type arpdhcp0151 struct{}

func Newarpdhcp0151() *arpdhcp0151 {
    return &arpdhcp0151{}
}

func (e *arpdhcp0151) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0151) Name() string { return "arpdhcp0151" }
func (e *arpdhcp0151) Timestamp() time.Time { return time.Now() }
