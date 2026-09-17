package passwordreset

import (
    "time"
)

type passwordreset0021 struct{}

func Newpasswordreset0021() *passwordreset0021 {
    return &passwordreset0021{}
}

func (e *passwordreset0021) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0021) Name() string { return "passwordreset0021" }
func (e *passwordreset0021) Timestamp() time.Time { return time.Now() }
