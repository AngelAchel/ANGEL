package passwordreset

import (
    "time"
)

type passwordreset0036 struct{}

func Newpasswordreset0036() *passwordreset0036 {
    return &passwordreset0036{}
}

func (e *passwordreset0036) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0036) Name() string { return "passwordreset0036" }
func (e *passwordreset0036) Timestamp() time.Time { return time.Now() }
