package passwordreset

import (
    "time"
)

type passwordreset0197 struct{}

func Newpasswordreset0197() *passwordreset0197 {
    return &passwordreset0197{}
}

func (e *passwordreset0197) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0197) Name() string { return "passwordreset0197" }
func (e *passwordreset0197) Timestamp() time.Time { return time.Now() }
