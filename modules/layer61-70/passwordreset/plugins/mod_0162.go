package passwordreset

import (
    "time"
)

type passwordreset0162 struct{}

func Newpasswordreset0162() *passwordreset0162 {
    return &passwordreset0162{}
}

func (e *passwordreset0162) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0162) Name() string { return "passwordreset0162" }
func (e *passwordreset0162) Timestamp() time.Time { return time.Now() }
