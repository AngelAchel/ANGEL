package arpdhcp

import (
    "time"
)

type arpdhcp0011 struct{}

func Newarpdhcp0011() *arpdhcp0011 {
    return &arpdhcp0011{}
}

func (e *arpdhcp0011) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0011) Name() string { return "arpdhcp0011" }
func (e *arpdhcp0011) Timestamp() time.Time { return time.Now() }
