package passwordreset

import (
    "time"
)

type passwordreset0150 struct{}

func Newpasswordreset0150() *passwordreset0150 {
    return &passwordreset0150{}
}

func (e *passwordreset0150) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0150) Name() string { return "passwordreset0150" }
func (e *passwordreset0150) Timestamp() time.Time { return time.Now() }
