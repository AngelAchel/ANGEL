package arpdhcp

import (
    "time"
)

type arpdhcp0059 struct{}

func Newarpdhcp0059() *arpdhcp0059 {
    return &arpdhcp0059{}
}

func (e *arpdhcp0059) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0059) Name() string { return "arpdhcp0059" }
func (e *arpdhcp0059) Timestamp() time.Time { return time.Now() }
