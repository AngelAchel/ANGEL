package passwordreset

import (
    "time"
)

type passwordreset0073 struct{}

func Newpasswordreset0073() *passwordreset0073 {
    return &passwordreset0073{}
}

func (e *passwordreset0073) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0073) Name() string { return "passwordreset0073" }
func (e *passwordreset0073) Timestamp() time.Time { return time.Now() }
