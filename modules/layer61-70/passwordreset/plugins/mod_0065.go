package passwordreset

import (
    "time"
)

type passwordreset0065 struct{}

func Newpasswordreset0065() *passwordreset0065 {
    return &passwordreset0065{}
}

func (e *passwordreset0065) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0065) Name() string { return "passwordreset0065" }
func (e *passwordreset0065) Timestamp() time.Time { return time.Now() }
