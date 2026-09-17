package passwordreset

import (
    "time"
)

type passwordreset0099 struct{}

func Newpasswordreset0099() *passwordreset0099 {
    return &passwordreset0099{}
}

func (e *passwordreset0099) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0099) Name() string { return "passwordreset0099" }
func (e *passwordreset0099) Timestamp() time.Time { return time.Now() }
