package memcorrupt

import (
    "time"
)

type memcorrupt0170 struct{}

func Newmemcorrupt0170() *memcorrupt0170 {
    return &memcorrupt0170{}
}

func (e *memcorrupt0170) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0170) Name() string { return "memcorrupt0170" }
func (e *memcorrupt0170) Timestamp() time.Time { return time.Now() }
