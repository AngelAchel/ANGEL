package passwordreset

import (
    "time"
)

type passwordreset0017 struct{}

func Newpasswordreset0017() *passwordreset0017 {
    return &passwordreset0017{}
}

func (e *passwordreset0017) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0017) Name() string { return "passwordreset0017" }
func (e *passwordreset0017) Timestamp() time.Time { return time.Now() }
