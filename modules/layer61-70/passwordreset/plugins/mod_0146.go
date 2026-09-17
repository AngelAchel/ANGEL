package passwordreset

import (
    "time"
)

type passwordreset0146 struct{}

func Newpasswordreset0146() *passwordreset0146 {
    return &passwordreset0146{}
}

func (e *passwordreset0146) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0146) Name() string { return "passwordreset0146" }
func (e *passwordreset0146) Timestamp() time.Time { return time.Now() }
