package passwordreset

import (
    "time"
)

type passwordreset0074 struct{}

func Newpasswordreset0074() *passwordreset0074 {
    return &passwordreset0074{}
}

func (e *passwordreset0074) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0074) Name() string { return "passwordreset0074" }
func (e *passwordreset0074) Timestamp() time.Time { return time.Now() }
