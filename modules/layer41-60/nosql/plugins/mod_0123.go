package nosql

import (
    "time"
)

type nosql0123 struct{}

func Newnosql0123() *nosql0123 {
    return &nosql0123{}
}

func (e *nosql0123) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "nosql:done")
    return results, nil
}

func (e *nosql0123) Name() string { return "nosql0123" }
func (e *nosql0123) Timestamp() time.Time { return time.Now() }
