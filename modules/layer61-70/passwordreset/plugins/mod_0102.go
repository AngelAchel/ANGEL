package passwordreset

import (
    "time"
)

type passwordreset0102 struct{}

func Newpasswordreset0102() *passwordreset0102 {
    return &passwordreset0102{}
}

func (e *passwordreset0102) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0102) Name() string { return "passwordreset0102" }
func (e *passwordreset0102) Timestamp() time.Time { return time.Now() }
