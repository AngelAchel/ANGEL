package passwordreset

import (
    "time"
)

type passwordreset0145 struct{}

func Newpasswordreset0145() *passwordreset0145 {
    return &passwordreset0145{}
}

func (e *passwordreset0145) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0145) Name() string { return "passwordreset0145" }
func (e *passwordreset0145) Timestamp() time.Time { return time.Now() }
