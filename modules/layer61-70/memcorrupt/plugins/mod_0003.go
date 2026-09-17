package memcorrupt

import (
    "time"
)

type memcorrupt0003 struct{}

func Newmemcorrupt0003() *memcorrupt0003 {
    return &memcorrupt0003{}
}

func (e *memcorrupt0003) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0003) Name() string { return "memcorrupt0003" }
func (e *memcorrupt0003) Timestamp() time.Time { return time.Now() }
