package deser

import (
    "time"
)

type deser0037 struct{}

func Newdeser0037() *deser0037 {
    return &deser0037{}
}

func (e *deser0037) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0037) Name() string { return "deser0037" }
func (e *deser0037) Timestamp() time.Time { return time.Now() }
