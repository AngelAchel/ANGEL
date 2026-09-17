package arpdhcp

import (
    "time"
)

type arpdhcp0106 struct{}

func Newarpdhcp0106() *arpdhcp0106 {
    return &arpdhcp0106{}
}

func (e *arpdhcp0106) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0106) Name() string { return "arpdhcp0106" }
func (e *arpdhcp0106) Timestamp() time.Time { return time.Now() }
