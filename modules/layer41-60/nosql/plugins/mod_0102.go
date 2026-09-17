package nosql

import (
    "time"
)

type nosql0102 struct{}

func Newnosql0102() *nosql0102 {
    return &nosql0102{}
}

func (e *nosql0102) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "nosql:done")
    return results, nil
}

func (e *nosql0102) Name() string { return "nosql0102" }
func (e *nosql0102) Timestamp() time.Time { return time.Now() }
