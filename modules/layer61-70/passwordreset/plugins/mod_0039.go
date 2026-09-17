package passwordreset

import (
    "time"
)

type passwordreset0039 struct{}

func Newpasswordreset0039() *passwordreset0039 {
    return &passwordreset0039{}
}

func (e *passwordreset0039) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0039) Name() string { return "passwordreset0039" }
func (e *passwordreset0039) Timestamp() time.Time { return time.Now() }
