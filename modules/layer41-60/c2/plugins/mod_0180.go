package c2

import (
    "time"
)

type c20180 struct{}

func Newc20180() *c20180 {
    return &c20180{}
}

func (e *c20180) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "c2:done")
    return results, nil
}

func (e *c20180) Name() string { return "c20180" }
func (e *c20180) Timestamp() time.Time { return time.Now() }
