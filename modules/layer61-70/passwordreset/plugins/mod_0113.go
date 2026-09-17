package passwordreset

import (
    "time"
)

type passwordreset0113 struct{}

func Newpasswordreset0113() *passwordreset0113 {
    return &passwordreset0113{}
}

func (e *passwordreset0113) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0113) Name() string { return "passwordreset0113" }
func (e *passwordreset0113) Timestamp() time.Time { return time.Now() }
