package passwordreset

import (
    "time"
)

type passwordreset0106 struct{}

func Newpasswordreset0106() *passwordreset0106 {
    return &passwordreset0106{}
}

func (e *passwordreset0106) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0106) Name() string { return "passwordreset0106" }
func (e *passwordreset0106) Timestamp() time.Time { return time.Now() }
