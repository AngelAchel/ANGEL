package memcorrupt

import (
    "time"
)

type memcorrupt0128 struct{}

func Newmemcorrupt0128() *memcorrupt0128 {
    return &memcorrupt0128{}
}

func (e *memcorrupt0128) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0128) Name() string { return "memcorrupt0128" }
func (e *memcorrupt0128) Timestamp() time.Time { return time.Now() }
