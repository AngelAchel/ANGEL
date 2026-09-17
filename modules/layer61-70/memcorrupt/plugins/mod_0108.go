package memcorrupt

import (
    "time"
)

type memcorrupt0108 struct{}

func Newmemcorrupt0108() *memcorrupt0108 {
    return &memcorrupt0108{}
}

func (e *memcorrupt0108) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0108) Name() string { return "memcorrupt0108" }
func (e *memcorrupt0108) Timestamp() time.Time { return time.Now() }
