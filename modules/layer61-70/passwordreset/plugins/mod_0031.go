package passwordreset

import (
    "time"
)

type passwordreset0031 struct{}

func Newpasswordreset0031() *passwordreset0031 {
    return &passwordreset0031{}
}

func (e *passwordreset0031) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0031) Name() string { return "passwordreset0031" }
func (e *passwordreset0031) Timestamp() time.Time { return time.Now() }
