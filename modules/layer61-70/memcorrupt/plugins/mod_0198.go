package memcorrupt

import (
    "time"
)

type memcorrupt0198 struct{}

func Newmemcorrupt0198() *memcorrupt0198 {
    return &memcorrupt0198{}
}

func (e *memcorrupt0198) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0198) Name() string { return "memcorrupt0198" }
func (e *memcorrupt0198) Timestamp() time.Time { return time.Now() }
