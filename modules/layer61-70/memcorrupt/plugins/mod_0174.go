package memcorrupt

import (
    "time"
)

type memcorrupt0174 struct{}

func Newmemcorrupt0174() *memcorrupt0174 {
    return &memcorrupt0174{}
}

func (e *memcorrupt0174) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0174) Name() string { return "memcorrupt0174" }
func (e *memcorrupt0174) Timestamp() time.Time { return time.Now() }
