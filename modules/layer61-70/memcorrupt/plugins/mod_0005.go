package memcorrupt

import (
    "time"
)

type memcorrupt0005 struct{}

func Newmemcorrupt0005() *memcorrupt0005 {
    return &memcorrupt0005{}
}

func (e *memcorrupt0005) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0005) Name() string { return "memcorrupt0005" }
func (e *memcorrupt0005) Timestamp() time.Time { return time.Now() }
