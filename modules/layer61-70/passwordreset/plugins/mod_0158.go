package passwordreset

import (
    "time"
)

type passwordreset0158 struct{}

func Newpasswordreset0158() *passwordreset0158 {
    return &passwordreset0158{}
}

func (e *passwordreset0158) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0158) Name() string { return "passwordreset0158" }
func (e *passwordreset0158) Timestamp() time.Time { return time.Now() }
