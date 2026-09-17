package passwordreset

import (
    "time"
)

type passwordreset0091 struct{}

func Newpasswordreset0091() *passwordreset0091 {
    return &passwordreset0091{}
}

func (e *passwordreset0091) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0091) Name() string { return "passwordreset0091" }
func (e *passwordreset0091) Timestamp() time.Time { return time.Now() }
