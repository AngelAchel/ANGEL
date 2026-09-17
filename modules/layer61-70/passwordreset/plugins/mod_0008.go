package passwordreset

import (
    "time"
)

type passwordreset0008 struct{}

func Newpasswordreset0008() *passwordreset0008 {
    return &passwordreset0008{}
}

func (e *passwordreset0008) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0008) Name() string { return "passwordreset0008" }
func (e *passwordreset0008) Timestamp() time.Time { return time.Now() }
