package memcorrupt

import (
    "time"
)

type memcorrupt0080 struct{}

func Newmemcorrupt0080() *memcorrupt0080 {
    return &memcorrupt0080{}
}

func (e *memcorrupt0080) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0080) Name() string { return "memcorrupt0080" }
func (e *memcorrupt0080) Timestamp() time.Time { return time.Now() }
