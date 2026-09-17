package memcorrupt

import (
    "time"
)

type memcorrupt0106 struct{}

func Newmemcorrupt0106() *memcorrupt0106 {
    return &memcorrupt0106{}
}

func (e *memcorrupt0106) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0106) Name() string { return "memcorrupt0106" }
func (e *memcorrupt0106) Timestamp() time.Time { return time.Now() }
