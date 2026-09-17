package webmisc

import (
	"time"
)

type UARotate struct{}

func NewUARotate() *UARotate {
	return &UARotate{}
}

func (u *UARotate) Rotate() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "ua_rotate:done")
	return results, nil
}

func (u *UARotate) Name() string { return "UARotate" }
func (u *UARotate) Timestamp() time.Time { return time.Now() }
