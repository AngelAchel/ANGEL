package passwordreset

import (
    "time"
)

type passwordreset0095 struct{}

func Newpasswordreset0095() *passwordreset0095 {
    return &passwordreset0095{}
}

func (e *passwordreset0095) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0095) Name() string { return "passwordreset0095" }
func (e *passwordreset0095) Timestamp() time.Time { return time.Now() }
