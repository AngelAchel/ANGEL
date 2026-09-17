package arpdhcp

import (
    "time"
)

type arpdhcp0029 struct{}

func Newarpdhcp0029() *arpdhcp0029 {
    return &arpdhcp0029{}
}

func (e *arpdhcp0029) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0029) Name() string { return "arpdhcp0029" }
func (e *arpdhcp0029) Timestamp() time.Time { return time.Now() }
