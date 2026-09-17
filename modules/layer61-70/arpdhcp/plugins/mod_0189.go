package arpdhcp

import (
    "time"
)

type arpdhcp0189 struct{}

func Newarpdhcp0189() *arpdhcp0189 {
    return &arpdhcp0189{}
}

func (e *arpdhcp0189) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0189) Name() string { return "arpdhcp0189" }
func (e *arpdhcp0189) Timestamp() time.Time { return time.Now() }
