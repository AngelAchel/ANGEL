package memcorrupt

import (
    "time"
)

type memcorrupt0138 struct{}

func Newmemcorrupt0138() *memcorrupt0138 {
    return &memcorrupt0138{}
}

func (e *memcorrupt0138) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0138) Name() string { return "memcorrupt0138" }
func (e *memcorrupt0138) Timestamp() time.Time { return time.Now() }
