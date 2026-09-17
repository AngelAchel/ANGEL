package memcorrupt

import (
    "time"
)

type memcorrupt0087 struct{}

func Newmemcorrupt0087() *memcorrupt0087 {
    return &memcorrupt0087{}
}

func (e *memcorrupt0087) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0087) Name() string { return "memcorrupt0087" }
func (e *memcorrupt0087) Timestamp() time.Time { return time.Now() }
