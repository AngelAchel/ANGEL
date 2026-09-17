package memcorrupt

import (
    "time"
)

type memcorrupt0136 struct{}

func Newmemcorrupt0136() *memcorrupt0136 {
    return &memcorrupt0136{}
}

func (e *memcorrupt0136) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0136) Name() string { return "memcorrupt0136" }
func (e *memcorrupt0136) Timestamp() time.Time { return time.Now() }
