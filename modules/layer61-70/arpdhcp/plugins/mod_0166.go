package arpdhcp

import (
    "time"
)

type arpdhcp0166 struct{}

func Newarpdhcp0166() *arpdhcp0166 {
    return &arpdhcp0166{}
}

func (e *arpdhcp0166) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0166) Name() string { return "arpdhcp0166" }
func (e *arpdhcp0166) Timestamp() time.Time { return time.Now() }
