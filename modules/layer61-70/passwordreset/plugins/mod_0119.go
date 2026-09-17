package passwordreset

import (
    "time"
)

type passwordreset0119 struct{}

func Newpasswordreset0119() *passwordreset0119 {
    return &passwordreset0119{}
}

func (e *passwordreset0119) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0119) Name() string { return "passwordreset0119" }
func (e *passwordreset0119) Timestamp() time.Time { return time.Now() }
