package passwordreset

import (
    "time"
)

type passwordreset0177 struct{}

func Newpasswordreset0177() *passwordreset0177 {
    return &passwordreset0177{}
}

func (e *passwordreset0177) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0177) Name() string { return "passwordreset0177" }
func (e *passwordreset0177) Timestamp() time.Time { return time.Now() }
