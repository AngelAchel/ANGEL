package passwordreset

import (
    "time"
)

type passwordreset0124 struct{}

func Newpasswordreset0124() *passwordreset0124 {
    return &passwordreset0124{}
}

func (e *passwordreset0124) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0124) Name() string { return "passwordreset0124" }
func (e *passwordreset0124) Timestamp() time.Time { return time.Now() }
