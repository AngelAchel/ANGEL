package passwordreset

import (
    "time"
)

type passwordreset0195 struct{}

func Newpasswordreset0195() *passwordreset0195 {
    return &passwordreset0195{}
}

func (e *passwordreset0195) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0195) Name() string { return "passwordreset0195" }
func (e *passwordreset0195) Timestamp() time.Time { return time.Now() }
