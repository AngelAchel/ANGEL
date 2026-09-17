package passwordreset

import (
    "time"
)

type passwordreset0191 struct{}

func Newpasswordreset0191() *passwordreset0191 {
    return &passwordreset0191{}
}

func (e *passwordreset0191) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0191) Name() string { return "passwordreset0191" }
func (e *passwordreset0191) Timestamp() time.Time { return time.Now() }
