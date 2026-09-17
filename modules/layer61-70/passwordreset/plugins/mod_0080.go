package passwordreset

import (
    "time"
)

type passwordreset0080 struct{}

func Newpasswordreset0080() *passwordreset0080 {
    return &passwordreset0080{}
}

func (e *passwordreset0080) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0080) Name() string { return "passwordreset0080" }
func (e *passwordreset0080) Timestamp() time.Time { return time.Now() }
