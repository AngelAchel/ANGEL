package arpdhcp

import (
    "time"
)

type arpdhcp0065 struct{}

func Newarpdhcp0065() *arpdhcp0065 {
    return &arpdhcp0065{}
}

func (e *arpdhcp0065) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0065) Name() string { return "arpdhcp0065" }
func (e *arpdhcp0065) Timestamp() time.Time { return time.Now() }
