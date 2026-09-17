package arpdhcp

import (
    "time"
)

type arpdhcp0027 struct{}

func Newarpdhcp0027() *arpdhcp0027 {
    return &arpdhcp0027{}
}

func (e *arpdhcp0027) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0027) Name() string { return "arpdhcp0027" }
func (e *arpdhcp0027) Timestamp() time.Time { return time.Now() }
