package memcorrupt

import (
    "time"
)

type memcorrupt0104 struct{}

func Newmemcorrupt0104() *memcorrupt0104 {
    return &memcorrupt0104{}
}

func (e *memcorrupt0104) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0104) Name() string { return "memcorrupt0104" }
func (e *memcorrupt0104) Timestamp() time.Time { return time.Now() }
