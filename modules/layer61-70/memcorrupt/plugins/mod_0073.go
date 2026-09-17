package memcorrupt

import (
    "time"
)

type memcorrupt0073 struct{}

func Newmemcorrupt0073() *memcorrupt0073 {
    return &memcorrupt0073{}
}

func (e *memcorrupt0073) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0073) Name() string { return "memcorrupt0073" }
func (e *memcorrupt0073) Timestamp() time.Time { return time.Now() }
