package arpdhcp

import (
    "time"
)

type arpdhcp0067 struct{}

func Newarpdhcp0067() *arpdhcp0067 {
    return &arpdhcp0067{}
}

func (e *arpdhcp0067) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0067) Name() string { return "arpdhcp0067" }
func (e *arpdhcp0067) Timestamp() time.Time { return time.Now() }
