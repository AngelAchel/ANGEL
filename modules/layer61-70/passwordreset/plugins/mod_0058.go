package passwordreset

import (
    "time"
)

type passwordreset0058 struct{}

func Newpasswordreset0058() *passwordreset0058 {
    return &passwordreset0058{}
}

func (e *passwordreset0058) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0058) Name() string { return "passwordreset0058" }
func (e *passwordreset0058) Timestamp() time.Time { return time.Now() }
