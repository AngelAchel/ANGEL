package memcorrupt

import (
    "time"
)

type memcorrupt0001 struct{}

func Newmemcorrupt0001() *memcorrupt0001 {
    return &memcorrupt0001{}
}

func (e *memcorrupt0001) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0001) Name() string { return "memcorrupt0001" }
func (e *memcorrupt0001) Timestamp() time.Time { return time.Now() }
