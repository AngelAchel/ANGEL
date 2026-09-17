package passwordreset

import (
    "time"
)

type passwordreset0004 struct{}

func Newpasswordreset0004() *passwordreset0004 {
    return &passwordreset0004{}
}

func (e *passwordreset0004) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0004) Name() string { return "passwordreset0004" }
func (e *passwordreset0004) Timestamp() time.Time { return time.Now() }
