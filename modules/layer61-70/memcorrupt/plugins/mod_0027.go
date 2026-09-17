package memcorrupt

import (
    "time"
)

type memcorrupt0027 struct{}

func Newmemcorrupt0027() *memcorrupt0027 {
    return &memcorrupt0027{}
}

func (e *memcorrupt0027) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0027) Name() string { return "memcorrupt0027" }
func (e *memcorrupt0027) Timestamp() time.Time { return time.Now() }
