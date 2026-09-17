package passwordreset

import (
    "time"
)

type passwordreset0096 struct{}

func Newpasswordreset0096() *passwordreset0096 {
    return &passwordreset0096{}
}

func (e *passwordreset0096) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0096) Name() string { return "passwordreset0096" }
func (e *passwordreset0096) Timestamp() time.Time { return time.Now() }
