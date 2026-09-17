package passwordreset

import (
    "time"
)

type passwordreset0116 struct{}

func Newpasswordreset0116() *passwordreset0116 {
    return &passwordreset0116{}
}

func (e *passwordreset0116) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0116) Name() string { return "passwordreset0116" }
func (e *passwordreset0116) Timestamp() time.Time { return time.Now() }
