package arpdhcp

import (
    "time"
)

type arpdhcp0056 struct{}

func Newarpdhcp0056() *arpdhcp0056 {
    return &arpdhcp0056{}
}

func (e *arpdhcp0056) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0056) Name() string { return "arpdhcp0056" }
func (e *arpdhcp0056) Timestamp() time.Time { return time.Now() }
