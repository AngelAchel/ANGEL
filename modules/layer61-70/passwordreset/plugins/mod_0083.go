package passwordreset

import (
    "time"
)

type passwordreset0083 struct{}

func Newpasswordreset0083() *passwordreset0083 {
    return &passwordreset0083{}
}

func (e *passwordreset0083) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0083) Name() string { return "passwordreset0083" }
func (e *passwordreset0083) Timestamp() time.Time { return time.Now() }
