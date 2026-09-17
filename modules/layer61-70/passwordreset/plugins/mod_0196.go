package passwordreset

import (
    "time"
)

type passwordreset0196 struct{}

func Newpasswordreset0196() *passwordreset0196 {
    return &passwordreset0196{}
}

func (e *passwordreset0196) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0196) Name() string { return "passwordreset0196" }
func (e *passwordreset0196) Timestamp() time.Time { return time.Now() }
