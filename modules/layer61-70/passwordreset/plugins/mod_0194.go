package passwordreset

import (
    "time"
)

type passwordreset0194 struct{}

func Newpasswordreset0194() *passwordreset0194 {
    return &passwordreset0194{}
}

func (e *passwordreset0194) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0194) Name() string { return "passwordreset0194" }
func (e *passwordreset0194) Timestamp() time.Time { return time.Now() }
