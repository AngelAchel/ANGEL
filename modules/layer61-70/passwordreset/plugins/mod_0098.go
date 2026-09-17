package passwordreset

import (
    "time"
)

type passwordreset0098 struct{}

func Newpasswordreset0098() *passwordreset0098 {
    return &passwordreset0098{}
}

func (e *passwordreset0098) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0098) Name() string { return "passwordreset0098" }
func (e *passwordreset0098) Timestamp() time.Time { return time.Now() }
