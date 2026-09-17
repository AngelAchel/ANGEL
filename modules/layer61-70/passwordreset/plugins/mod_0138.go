package passwordreset

import (
    "time"
)

type passwordreset0138 struct{}

func Newpasswordreset0138() *passwordreset0138 {
    return &passwordreset0138{}
}

func (e *passwordreset0138) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0138) Name() string { return "passwordreset0138" }
func (e *passwordreset0138) Timestamp() time.Time { return time.Now() }
