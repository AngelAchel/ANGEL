package passwordreset

import (
    "time"
)

type passwordreset0187 struct{}

func Newpasswordreset0187() *passwordreset0187 {
    return &passwordreset0187{}
}

func (e *passwordreset0187) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0187) Name() string { return "passwordreset0187" }
func (e *passwordreset0187) Timestamp() time.Time { return time.Now() }
