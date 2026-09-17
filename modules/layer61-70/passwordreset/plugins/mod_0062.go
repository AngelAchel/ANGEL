package passwordreset

import (
    "time"
)

type passwordreset0062 struct{}

func Newpasswordreset0062() *passwordreset0062 {
    return &passwordreset0062{}
}

func (e *passwordreset0062) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0062) Name() string { return "passwordreset0062" }
func (e *passwordreset0062) Timestamp() time.Time { return time.Now() }
