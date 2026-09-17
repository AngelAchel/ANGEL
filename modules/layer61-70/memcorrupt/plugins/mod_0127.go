package memcorrupt

import (
    "time"
)

type memcorrupt0127 struct{}

func Newmemcorrupt0127() *memcorrupt0127 {
    return &memcorrupt0127{}
}

func (e *memcorrupt0127) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0127) Name() string { return "memcorrupt0127" }
func (e *memcorrupt0127) Timestamp() time.Time { return time.Now() }
