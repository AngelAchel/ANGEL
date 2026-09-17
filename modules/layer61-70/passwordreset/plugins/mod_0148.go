package passwordreset

import (
    "time"
)

type passwordreset0148 struct{}

func Newpasswordreset0148() *passwordreset0148 {
    return &passwordreset0148{}
}

func (e *passwordreset0148) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0148) Name() string { return "passwordreset0148" }
func (e *passwordreset0148) Timestamp() time.Time { return time.Now() }
