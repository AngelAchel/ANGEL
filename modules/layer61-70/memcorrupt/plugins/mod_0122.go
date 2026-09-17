package memcorrupt

import (
    "time"
)

type memcorrupt0122 struct{}

func Newmemcorrupt0122() *memcorrupt0122 {
    return &memcorrupt0122{}
}

func (e *memcorrupt0122) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0122) Name() string { return "memcorrupt0122" }
func (e *memcorrupt0122) Timestamp() time.Time { return time.Now() }
