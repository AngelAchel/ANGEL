package passwordreset

import (
    "time"
)

type passwordreset0171 struct{}

func Newpasswordreset0171() *passwordreset0171 {
    return &passwordreset0171{}
}

func (e *passwordreset0171) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0171) Name() string { return "passwordreset0171" }
func (e *passwordreset0171) Timestamp() time.Time { return time.Now() }
