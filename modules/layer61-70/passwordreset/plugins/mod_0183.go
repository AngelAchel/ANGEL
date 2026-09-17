package passwordreset

import (
    "time"
)

type passwordreset0183 struct{}

func Newpasswordreset0183() *passwordreset0183 {
    return &passwordreset0183{}
}

func (e *passwordreset0183) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0183) Name() string { return "passwordreset0183" }
func (e *passwordreset0183) Timestamp() time.Time { return time.Now() }
