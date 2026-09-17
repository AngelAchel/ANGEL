package passwordreset

import (
    "time"
)

type passwordreset0001 struct{}

func Newpasswordreset0001() *passwordreset0001 {
    return &passwordreset0001{}
}

func (e *passwordreset0001) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0001) Name() string { return "passwordreset0001" }
func (e *passwordreset0001) Timestamp() time.Time { return time.Now() }
