package memcorrupt

import (
    "time"
)

type memcorrupt0134 struct{}

func Newmemcorrupt0134() *memcorrupt0134 {
    return &memcorrupt0134{}
}

func (e *memcorrupt0134) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0134) Name() string { return "memcorrupt0134" }
func (e *memcorrupt0134) Timestamp() time.Time { return time.Now() }
