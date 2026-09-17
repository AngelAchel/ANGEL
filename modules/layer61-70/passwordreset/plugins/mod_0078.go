package passwordreset

import (
    "time"
)

type passwordreset0078 struct{}

func Newpasswordreset0078() *passwordreset0078 {
    return &passwordreset0078{}
}

func (e *passwordreset0078) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0078) Name() string { return "passwordreset0078" }
func (e *passwordreset0078) Timestamp() time.Time { return time.Now() }
