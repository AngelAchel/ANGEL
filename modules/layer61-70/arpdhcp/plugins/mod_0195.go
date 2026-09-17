package arpdhcp

import (
    "time"
)

type arpdhcp0195 struct{}

func Newarpdhcp0195() *arpdhcp0195 {
    return &arpdhcp0195{}
}

func (e *arpdhcp0195) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0195) Name() string { return "arpdhcp0195" }
func (e *arpdhcp0195) Timestamp() time.Time { return time.Now() }
