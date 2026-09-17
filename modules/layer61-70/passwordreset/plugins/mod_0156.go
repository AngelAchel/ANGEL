package passwordreset

import (
    "time"
)

type passwordreset0156 struct{}

func Newpasswordreset0156() *passwordreset0156 {
    return &passwordreset0156{}
}

func (e *passwordreset0156) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0156) Name() string { return "passwordreset0156" }
func (e *passwordreset0156) Timestamp() time.Time { return time.Now() }
