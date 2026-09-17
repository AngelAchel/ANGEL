package nosql

import (
    "time"
)

type nosql0100 struct{}

func Newnosql0100() *nosql0100 {
    return &nosql0100{}
}

func (e *nosql0100) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "nosql:done")
    return results, nil
}

func (e *nosql0100) Name() string { return "nosql0100" }
func (e *nosql0100) Timestamp() time.Time { return time.Now() }
