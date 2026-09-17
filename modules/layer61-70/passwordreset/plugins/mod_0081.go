package passwordreset

import (
    "time"
)

type passwordreset0081 struct{}

func Newpasswordreset0081() *passwordreset0081 {
    return &passwordreset0081{}
}

func (e *passwordreset0081) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0081) Name() string { return "passwordreset0081" }
func (e *passwordreset0081) Timestamp() time.Time { return time.Now() }
