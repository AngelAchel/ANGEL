package deser

import (
    "time"
)

type deser0003 struct{}

func Newdeser0003() *deser0003 {
    return &deser0003{}
}

func (e *deser0003) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0003) Name() string { return "deser0003" }
func (e *deser0003) Timestamp() time.Time { return time.Now() }
