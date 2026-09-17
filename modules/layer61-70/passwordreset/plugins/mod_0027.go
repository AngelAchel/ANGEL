package passwordreset

import (
    "time"
)

type passwordreset0027 struct{}

func Newpasswordreset0027() *passwordreset0027 {
    return &passwordreset0027{}
}

func (e *passwordreset0027) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0027) Name() string { return "passwordreset0027" }
func (e *passwordreset0027) Timestamp() time.Time { return time.Now() }
