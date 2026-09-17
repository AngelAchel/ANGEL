package memcorrupt

import (
    "time"
)

type memcorrupt0102 struct{}

func Newmemcorrupt0102() *memcorrupt0102 {
    return &memcorrupt0102{}
}

func (e *memcorrupt0102) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0102) Name() string { return "memcorrupt0102" }
func (e *memcorrupt0102) Timestamp() time.Time { return time.Now() }
