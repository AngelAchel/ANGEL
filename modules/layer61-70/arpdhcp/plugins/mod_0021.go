package arpdhcp

import (
    "time"
)

type arpdhcp0021 struct{}

func Newarpdhcp0021() *arpdhcp0021 {
    return &arpdhcp0021{}
}

func (e *arpdhcp0021) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0021) Name() string { return "arpdhcp0021" }
func (e *arpdhcp0021) Timestamp() time.Time { return time.Now() }
