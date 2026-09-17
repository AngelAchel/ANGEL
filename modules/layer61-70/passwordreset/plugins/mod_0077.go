package passwordreset

import (
    "time"
)

type passwordreset0077 struct{}

func Newpasswordreset0077() *passwordreset0077 {
    return &passwordreset0077{}
}

func (e *passwordreset0077) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0077) Name() string { return "passwordreset0077" }
func (e *passwordreset0077) Timestamp() time.Time { return time.Now() }
