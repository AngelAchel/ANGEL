package memcorrupt

import (
    "time"
)

type memcorrupt0111 struct{}

func Newmemcorrupt0111() *memcorrupt0111 {
    return &memcorrupt0111{}
}

func (e *memcorrupt0111) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0111) Name() string { return "memcorrupt0111" }
func (e *memcorrupt0111) Timestamp() time.Time { return time.Now() }
