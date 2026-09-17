package memcorrupt

import (
    "time"
)

type memcorrupt0081 struct{}

func Newmemcorrupt0081() *memcorrupt0081 {
    return &memcorrupt0081{}
}

func (e *memcorrupt0081) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0081) Name() string { return "memcorrupt0081" }
func (e *memcorrupt0081) Timestamp() time.Time { return time.Now() }
