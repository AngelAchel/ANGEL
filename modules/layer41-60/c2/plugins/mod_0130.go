package c2

import (
    "time"
)

type c20130 struct{}

func Newc20130() *c20130 {
    return &c20130{}
}

func (e *c20130) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "c2:done")
    return results, nil
}

func (e *c20130) Name() string { return "c20130" }
func (e *c20130) Timestamp() time.Time { return time.Now() }
