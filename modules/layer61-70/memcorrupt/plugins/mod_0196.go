package memcorrupt

import (
    "time"
)

type memcorrupt0196 struct{}

func Newmemcorrupt0196() *memcorrupt0196 {
    return &memcorrupt0196{}
}

func (e *memcorrupt0196) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0196) Name() string { return "memcorrupt0196" }
func (e *memcorrupt0196) Timestamp() time.Time { return time.Now() }
