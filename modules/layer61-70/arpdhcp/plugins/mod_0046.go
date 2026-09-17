package arpdhcp

import (
    "time"
)

type arpdhcp0046 struct{}

func Newarpdhcp0046() *arpdhcp0046 {
    return &arpdhcp0046{}
}

func (e *arpdhcp0046) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0046) Name() string { return "arpdhcp0046" }
func (e *arpdhcp0046) Timestamp() time.Time { return time.Now() }
