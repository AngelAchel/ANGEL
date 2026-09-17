package passwordreset

import (
    "time"
)

type passwordreset0151 struct{}

func Newpasswordreset0151() *passwordreset0151 {
    return &passwordreset0151{}
}

func (e *passwordreset0151) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0151) Name() string { return "passwordreset0151" }
func (e *passwordreset0151) Timestamp() time.Time { return time.Now() }
