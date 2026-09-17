package passwordreset

import (
    "time"
)

type passwordreset0064 struct{}

func Newpasswordreset0064() *passwordreset0064 {
    return &passwordreset0064{}
}

func (e *passwordreset0064) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0064) Name() string { return "passwordreset0064" }
func (e *passwordreset0064) Timestamp() time.Time { return time.Now() }
