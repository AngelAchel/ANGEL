package passwordreset

import (
    "time"
)

type passwordreset0032 struct{}

func Newpasswordreset0032() *passwordreset0032 {
    return &passwordreset0032{}
}

func (e *passwordreset0032) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0032) Name() string { return "passwordreset0032" }
func (e *passwordreset0032) Timestamp() time.Time { return time.Now() }
