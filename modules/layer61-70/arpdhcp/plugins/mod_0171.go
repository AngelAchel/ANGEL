package arpdhcp

import (
    "time"
)

type arpdhcp0171 struct{}

func Newarpdhcp0171() *arpdhcp0171 {
    return &arpdhcp0171{}
}

func (e *arpdhcp0171) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0171) Name() string { return "arpdhcp0171" }
func (e *arpdhcp0171) Timestamp() time.Time { return time.Now() }
