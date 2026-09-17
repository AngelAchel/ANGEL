package passwordreset

import (
    "time"
)

type passwordreset0184 struct{}

func Newpasswordreset0184() *passwordreset0184 {
    return &passwordreset0184{}
}

func (e *passwordreset0184) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0184) Name() string { return "passwordreset0184" }
func (e *passwordreset0184) Timestamp() time.Time { return time.Now() }
