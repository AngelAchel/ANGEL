package memcorrupt

import (
    "time"
)

type memcorrupt0097 struct{}

func Newmemcorrupt0097() *memcorrupt0097 {
    return &memcorrupt0097{}
}

func (e *memcorrupt0097) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0097) Name() string { return "memcorrupt0097" }
func (e *memcorrupt0097) Timestamp() time.Time { return time.Now() }
