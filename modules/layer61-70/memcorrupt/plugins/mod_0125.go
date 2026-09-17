package memcorrupt

import (
    "time"
)

type memcorrupt0125 struct{}

func Newmemcorrupt0125() *memcorrupt0125 {
    return &memcorrupt0125{}
}

func (e *memcorrupt0125) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0125) Name() string { return "memcorrupt0125" }
func (e *memcorrupt0125) Timestamp() time.Time { return time.Now() }
