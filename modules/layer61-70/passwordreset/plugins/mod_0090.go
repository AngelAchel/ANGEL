package passwordreset

import (
    "time"
)

type passwordreset0090 struct{}

func Newpasswordreset0090() *passwordreset0090 {
    return &passwordreset0090{}
}

func (e *passwordreset0090) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0090) Name() string { return "passwordreset0090" }
func (e *passwordreset0090) Timestamp() time.Time { return time.Now() }
