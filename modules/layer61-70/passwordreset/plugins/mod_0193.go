package passwordreset

import (
    "time"
)

type passwordreset0193 struct{}

func Newpasswordreset0193() *passwordreset0193 {
    return &passwordreset0193{}
}

func (e *passwordreset0193) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0193) Name() string { return "passwordreset0193" }
func (e *passwordreset0193) Timestamp() time.Time { return time.Now() }
