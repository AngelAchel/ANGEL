package memcorrupt

import (
    "time"
)

type memcorrupt0002 struct{}

func Newmemcorrupt0002() *memcorrupt0002 {
    return &memcorrupt0002{}
}

func (e *memcorrupt0002) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0002) Name() string { return "memcorrupt0002" }
func (e *memcorrupt0002) Timestamp() time.Time { return time.Now() }
