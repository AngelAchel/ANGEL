package memcorrupt

import (
    "time"
)

type memcorrupt0098 struct{}

func Newmemcorrupt0098() *memcorrupt0098 {
    return &memcorrupt0098{}
}

func (e *memcorrupt0098) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0098) Name() string { return "memcorrupt0098" }
func (e *memcorrupt0098) Timestamp() time.Time { return time.Now() }
