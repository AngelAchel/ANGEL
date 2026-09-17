package passwordreset

import (
    "time"
)

type passwordreset0128 struct{}

func Newpasswordreset0128() *passwordreset0128 {
    return &passwordreset0128{}
}

func (e *passwordreset0128) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0128) Name() string { return "passwordreset0128" }
func (e *passwordreset0128) Timestamp() time.Time { return time.Now() }
