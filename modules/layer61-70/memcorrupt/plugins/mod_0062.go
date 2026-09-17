package memcorrupt

import (
    "time"
)

type memcorrupt0062 struct{}

func Newmemcorrupt0062() *memcorrupt0062 {
    return &memcorrupt0062{}
}

func (e *memcorrupt0062) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0062) Name() string { return "memcorrupt0062" }
func (e *memcorrupt0062) Timestamp() time.Time { return time.Now() }
