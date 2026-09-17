package passwordreset

import (
    "time"
)

type passwordreset0122 struct{}

func Newpasswordreset0122() *passwordreset0122 {
    return &passwordreset0122{}
}

func (e *passwordreset0122) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0122) Name() string { return "passwordreset0122" }
func (e *passwordreset0122) Timestamp() time.Time { return time.Now() }
