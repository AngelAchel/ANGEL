package memcorrupt

import (
    "time"
)

type memcorrupt0130 struct{}

func Newmemcorrupt0130() *memcorrupt0130 {
    return &memcorrupt0130{}
}

func (e *memcorrupt0130) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0130) Name() string { return "memcorrupt0130" }
func (e *memcorrupt0130) Timestamp() time.Time { return time.Now() }
