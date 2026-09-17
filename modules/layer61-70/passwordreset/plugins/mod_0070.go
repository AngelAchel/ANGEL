package passwordreset

import (
    "time"
)

type passwordreset0070 struct{}

func Newpasswordreset0070() *passwordreset0070 {
    return &passwordreset0070{}
}

func (e *passwordreset0070) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0070) Name() string { return "passwordreset0070" }
func (e *passwordreset0070) Timestamp() time.Time { return time.Now() }
