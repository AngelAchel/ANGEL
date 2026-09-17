package passwordreset

import (
    "time"
)

type passwordreset0082 struct{}

func Newpasswordreset0082() *passwordreset0082 {
    return &passwordreset0082{}
}

func (e *passwordreset0082) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0082) Name() string { return "passwordreset0082" }
func (e *passwordreset0082) Timestamp() time.Time { return time.Now() }
