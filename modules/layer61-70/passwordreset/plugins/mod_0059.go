package passwordreset

import (
    "time"
)

type passwordreset0059 struct{}

func Newpasswordreset0059() *passwordreset0059 {
    return &passwordreset0059{}
}

func (e *passwordreset0059) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0059) Name() string { return "passwordreset0059" }
func (e *passwordreset0059) Timestamp() time.Time { return time.Now() }
