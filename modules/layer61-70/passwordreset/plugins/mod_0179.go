package passwordreset

import (
    "time"
)

type passwordreset0179 struct{}

func Newpasswordreset0179() *passwordreset0179 {
    return &passwordreset0179{}
}

func (e *passwordreset0179) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0179) Name() string { return "passwordreset0179" }
func (e *passwordreset0179) Timestamp() time.Time { return time.Now() }
