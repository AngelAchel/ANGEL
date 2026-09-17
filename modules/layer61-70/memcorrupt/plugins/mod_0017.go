package memcorrupt

import (
    "time"
)

type memcorrupt0017 struct{}

func Newmemcorrupt0017() *memcorrupt0017 {
    return &memcorrupt0017{}
}

func (e *memcorrupt0017) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0017) Name() string { return "memcorrupt0017" }
func (e *memcorrupt0017) Timestamp() time.Time { return time.Now() }
