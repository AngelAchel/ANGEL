package passwordreset

import (
    "time"
)

type passwordreset0007 struct{}

func Newpasswordreset0007() *passwordreset0007 {
    return &passwordreset0007{}
}

func (e *passwordreset0007) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0007) Name() string { return "passwordreset0007" }
func (e *passwordreset0007) Timestamp() time.Time { return time.Now() }
