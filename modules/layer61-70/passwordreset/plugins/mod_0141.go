package passwordreset

import (
    "time"
)

type passwordreset0141 struct{}

func Newpasswordreset0141() *passwordreset0141 {
    return &passwordreset0141{}
}

func (e *passwordreset0141) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0141) Name() string { return "passwordreset0141" }
func (e *passwordreset0141) Timestamp() time.Time { return time.Now() }
