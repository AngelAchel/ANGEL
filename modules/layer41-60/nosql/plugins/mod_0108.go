package nosql

import (
    "time"
)

type nosql0108 struct{}

func Newnosql0108() *nosql0108 {
    return &nosql0108{}
}

func (e *nosql0108) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "nosql:done")
    return results, nil
}

func (e *nosql0108) Name() string { return "nosql0108" }
func (e *nosql0108) Timestamp() time.Time { return time.Now() }
