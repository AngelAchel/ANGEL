package passwordreset

import (
    "time"
)

type passwordreset0176 struct{}

func Newpasswordreset0176() *passwordreset0176 {
    return &passwordreset0176{}
}

func (e *passwordreset0176) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0176) Name() string { return "passwordreset0176" }
func (e *passwordreset0176) Timestamp() time.Time { return time.Now() }
