package passwordreset

import (
    "time"
)

type passwordreset0103 struct{}

func Newpasswordreset0103() *passwordreset0103 {
    return &passwordreset0103{}
}

func (e *passwordreset0103) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0103) Name() string { return "passwordreset0103" }
func (e *passwordreset0103) Timestamp() time.Time { return time.Now() }
