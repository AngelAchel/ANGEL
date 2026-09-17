package memcorrupt

import (
    "time"
)

type memcorrupt0012 struct{}

func Newmemcorrupt0012() *memcorrupt0012 {
    return &memcorrupt0012{}
}

func (e *memcorrupt0012) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0012) Name() string { return "memcorrupt0012" }
func (e *memcorrupt0012) Timestamp() time.Time { return time.Now() }
