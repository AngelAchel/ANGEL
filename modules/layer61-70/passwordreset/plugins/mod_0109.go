package passwordreset

import (
    "time"
)

type passwordreset0109 struct{}

func Newpasswordreset0109() *passwordreset0109 {
    return &passwordreset0109{}
}

func (e *passwordreset0109) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0109) Name() string { return "passwordreset0109" }
func (e *passwordreset0109) Timestamp() time.Time { return time.Now() }
