package passwordreset

import (
    "time"
)

type passwordreset0046 struct{}

func Newpasswordreset0046() *passwordreset0046 {
    return &passwordreset0046{}
}

func (e *passwordreset0046) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0046) Name() string { return "passwordreset0046" }
func (e *passwordreset0046) Timestamp() time.Time { return time.Now() }
