package memcorrupt

import (
    "time"
)

type memcorrupt0038 struct{}

func Newmemcorrupt0038() *memcorrupt0038 {
    return &memcorrupt0038{}
}

func (e *memcorrupt0038) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0038) Name() string { return "memcorrupt0038" }
func (e *memcorrupt0038) Timestamp() time.Time { return time.Now() }
