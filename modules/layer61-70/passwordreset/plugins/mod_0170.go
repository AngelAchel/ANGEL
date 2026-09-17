package passwordreset

import (
    "time"
)

type passwordreset0170 struct{}

func Newpasswordreset0170() *passwordreset0170 {
    return &passwordreset0170{}
}

func (e *passwordreset0170) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0170) Name() string { return "passwordreset0170" }
func (e *passwordreset0170) Timestamp() time.Time { return time.Now() }
