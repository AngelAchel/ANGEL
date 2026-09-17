package arpdhcp

import (
    "time"
)

type arpdhcp0062 struct{}

func Newarpdhcp0062() *arpdhcp0062 {
    return &arpdhcp0062{}
}

func (e *arpdhcp0062) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0062) Name() string { return "arpdhcp0062" }
func (e *arpdhcp0062) Timestamp() time.Time { return time.Now() }
