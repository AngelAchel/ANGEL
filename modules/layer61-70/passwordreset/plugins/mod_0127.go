package passwordreset

import (
    "time"
)

type passwordreset0127 struct{}

func Newpasswordreset0127() *passwordreset0127 {
    return &passwordreset0127{}
}

func (e *passwordreset0127) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0127) Name() string { return "passwordreset0127" }
func (e *passwordreset0127) Timestamp() time.Time { return time.Now() }
