package nosql

import (
    "time"
)

type nosql0130 struct{}

func Newnosql0130() *nosql0130 {
    return &nosql0130{}
}

func (e *nosql0130) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "nosql:done")
    return results, nil
}

func (e *nosql0130) Name() string { return "nosql0130" }
func (e *nosql0130) Timestamp() time.Time { return time.Now() }
