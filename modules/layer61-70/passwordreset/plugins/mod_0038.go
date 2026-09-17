package passwordreset

import (
    "time"
)

type passwordreset0038 struct{}

func Newpasswordreset0038() *passwordreset0038 {
    return &passwordreset0038{}
}

func (e *passwordreset0038) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0038) Name() string { return "passwordreset0038" }
func (e *passwordreset0038) Timestamp() time.Time { return time.Now() }
