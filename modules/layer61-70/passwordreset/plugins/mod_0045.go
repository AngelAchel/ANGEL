package passwordreset

import (
    "time"
)

type passwordreset0045 struct{}

func Newpasswordreset0045() *passwordreset0045 {
    return &passwordreset0045{}
}

func (e *passwordreset0045) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0045) Name() string { return "passwordreset0045" }
func (e *passwordreset0045) Timestamp() time.Time { return time.Now() }
