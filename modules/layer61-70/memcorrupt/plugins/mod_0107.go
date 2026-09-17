package memcorrupt

import (
    "time"
)

type memcorrupt0107 struct{}

func Newmemcorrupt0107() *memcorrupt0107 {
    return &memcorrupt0107{}
}

func (e *memcorrupt0107) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0107) Name() string { return "memcorrupt0107" }
func (e *memcorrupt0107) Timestamp() time.Time { return time.Now() }
