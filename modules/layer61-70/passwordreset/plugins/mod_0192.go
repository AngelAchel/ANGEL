package passwordreset

import (
    "time"
)

type passwordreset0192 struct{}

func Newpasswordreset0192() *passwordreset0192 {
    return &passwordreset0192{}
}

func (e *passwordreset0192) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0192) Name() string { return "passwordreset0192" }
func (e *passwordreset0192) Timestamp() time.Time { return time.Now() }
