package passwordreset

import (
    "time"
)

type passwordreset0084 struct{}

func Newpasswordreset0084() *passwordreset0084 {
    return &passwordreset0084{}
}

func (e *passwordreset0084) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0084) Name() string { return "passwordreset0084" }
func (e *passwordreset0084) Timestamp() time.Time { return time.Now() }
