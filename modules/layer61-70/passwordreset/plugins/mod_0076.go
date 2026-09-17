package passwordreset

import (
    "time"
)

type passwordreset0076 struct{}

func Newpasswordreset0076() *passwordreset0076 {
    return &passwordreset0076{}
}

func (e *passwordreset0076) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0076) Name() string { return "passwordreset0076" }
func (e *passwordreset0076) Timestamp() time.Time { return time.Now() }
