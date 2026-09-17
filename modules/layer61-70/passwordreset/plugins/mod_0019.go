package passwordreset

import (
    "time"
)

type passwordreset0019 struct{}

func Newpasswordreset0019() *passwordreset0019 {
    return &passwordreset0019{}
}

func (e *passwordreset0019) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0019) Name() string { return "passwordreset0019" }
func (e *passwordreset0019) Timestamp() time.Time { return time.Now() }
