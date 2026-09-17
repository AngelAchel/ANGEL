package passwordreset

import (
    "time"
)

type passwordreset0181 struct{}

func Newpasswordreset0181() *passwordreset0181 {
    return &passwordreset0181{}
}

func (e *passwordreset0181) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0181) Name() string { return "passwordreset0181" }
func (e *passwordreset0181) Timestamp() time.Time { return time.Now() }
