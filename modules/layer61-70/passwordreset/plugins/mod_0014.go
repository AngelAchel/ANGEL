package passwordreset

import (
    "time"
)

type passwordreset0014 struct{}

func Newpasswordreset0014() *passwordreset0014 {
    return &passwordreset0014{}
}

func (e *passwordreset0014) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0014) Name() string { return "passwordreset0014" }
func (e *passwordreset0014) Timestamp() time.Time { return time.Now() }
