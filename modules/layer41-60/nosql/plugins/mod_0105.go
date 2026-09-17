package nosql

import (
    "time"
)

type nosql0105 struct{}

func Newnosql0105() *nosql0105 {
    return &nosql0105{}
}

func (e *nosql0105) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "nosql:done")
    return results, nil
}

func (e *nosql0105) Name() string { return "nosql0105" }
func (e *nosql0105) Timestamp() time.Time { return time.Now() }
