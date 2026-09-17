package arpdhcp

import (
    "time"
)

type arpdhcp0141 struct{}

func Newarpdhcp0141() *arpdhcp0141 {
    return &arpdhcp0141{}
}

func (e *arpdhcp0141) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0141) Name() string { return "arpdhcp0141" }
func (e *arpdhcp0141) Timestamp() time.Time { return time.Now() }
