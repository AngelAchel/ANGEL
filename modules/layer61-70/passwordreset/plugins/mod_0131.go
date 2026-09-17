package passwordreset

import (
    "time"
)

type passwordreset0131 struct{}

func Newpasswordreset0131() *passwordreset0131 {
    return &passwordreset0131{}
}

func (e *passwordreset0131) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0131) Name() string { return "passwordreset0131" }
func (e *passwordreset0131) Timestamp() time.Time { return time.Now() }
