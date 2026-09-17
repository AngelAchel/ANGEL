package passwordreset

import (
    "time"
)

type passwordreset0168 struct{}

func Newpasswordreset0168() *passwordreset0168 {
    return &passwordreset0168{}
}

func (e *passwordreset0168) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0168) Name() string { return "passwordreset0168" }
func (e *passwordreset0168) Timestamp() time.Time { return time.Now() }
