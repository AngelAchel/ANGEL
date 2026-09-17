package passwordreset

import (
    "time"
)

type passwordreset0100 struct{}

func Newpasswordreset0100() *passwordreset0100 {
    return &passwordreset0100{}
}

func (e *passwordreset0100) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0100) Name() string { return "passwordreset0100" }
func (e *passwordreset0100) Timestamp() time.Time { return time.Now() }
