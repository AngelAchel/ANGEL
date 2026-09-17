package memcorrupt

import (
    "time"
)

type memcorrupt0193 struct{}

func Newmemcorrupt0193() *memcorrupt0193 {
    return &memcorrupt0193{}
}

func (e *memcorrupt0193) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0193) Name() string { return "memcorrupt0193" }
func (e *memcorrupt0193) Timestamp() time.Time { return time.Now() }
