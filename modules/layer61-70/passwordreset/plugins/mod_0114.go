package passwordreset

import (
    "time"
)

type passwordreset0114 struct{}

func Newpasswordreset0114() *passwordreset0114 {
    return &passwordreset0114{}
}

func (e *passwordreset0114) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0114) Name() string { return "passwordreset0114" }
func (e *passwordreset0114) Timestamp() time.Time { return time.Now() }
