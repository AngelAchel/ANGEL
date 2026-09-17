package passwordreset

import (
    "time"
)

type passwordreset0063 struct{}

func Newpasswordreset0063() *passwordreset0063 {
    return &passwordreset0063{}
}

func (e *passwordreset0063) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0063) Name() string { return "passwordreset0063" }
func (e *passwordreset0063) Timestamp() time.Time { return time.Now() }
