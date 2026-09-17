package passwordreset

import (
    "time"
)

type passwordreset0047 struct{}

func Newpasswordreset0047() *passwordreset0047 {
    return &passwordreset0047{}
}

func (e *passwordreset0047) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0047) Name() string { return "passwordreset0047" }
func (e *passwordreset0047) Timestamp() time.Time { return time.Now() }
