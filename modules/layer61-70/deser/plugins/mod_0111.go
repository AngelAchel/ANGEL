package deser

import (
    "time"
)

type deser0111 struct{}

func Newdeser0111() *deser0111 {
    return &deser0111{}
}

func (e *deser0111) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0111) Name() string { return "deser0111" }
func (e *deser0111) Timestamp() time.Time { return time.Now() }
