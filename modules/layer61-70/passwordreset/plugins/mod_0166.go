package passwordreset

import (
    "time"
)

type passwordreset0166 struct{}

func Newpasswordreset0166() *passwordreset0166 {
    return &passwordreset0166{}
}

func (e *passwordreset0166) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0166) Name() string { return "passwordreset0166" }
func (e *passwordreset0166) Timestamp() time.Time { return time.Now() }
