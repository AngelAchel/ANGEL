package passwordreset

import (
    "time"
)

type passwordreset0005 struct{}

func Newpasswordreset0005() *passwordreset0005 {
    return &passwordreset0005{}
}

func (e *passwordreset0005) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0005) Name() string { return "passwordreset0005" }
func (e *passwordreset0005) Timestamp() time.Time { return time.Now() }
