package passwordreset

import (
    "time"
)

type passwordreset0075 struct{}

func Newpasswordreset0075() *passwordreset0075 {
    return &passwordreset0075{}
}

func (e *passwordreset0075) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0075) Name() string { return "passwordreset0075" }
func (e *passwordreset0075) Timestamp() time.Time { return time.Now() }
