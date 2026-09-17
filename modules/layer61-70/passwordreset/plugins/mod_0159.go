package passwordreset

import (
    "time"
)

type passwordreset0159 struct{}

func Newpasswordreset0159() *passwordreset0159 {
    return &passwordreset0159{}
}

func (e *passwordreset0159) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0159) Name() string { return "passwordreset0159" }
func (e *passwordreset0159) Timestamp() time.Time { return time.Now() }
