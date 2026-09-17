package passwordreset

import (
    "time"
)

type passwordreset0000 struct{}

func Newpasswordreset0000() *passwordreset0000 {
    return &passwordreset0000{}
}

func (e *passwordreset0000) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0000) Name() string { return "passwordreset0000" }
func (e *passwordreset0000) Timestamp() time.Time { return time.Now() }
