package c2

import (
	"time"
)

type OOBHTTP struct{}

func NewOOBHTTP() *OOBHTTP {
	return &OOBHTTP{}
}

func (o *OOBHTTP) Callback() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "oob_http:done")
	return results, nil
}

func (o *OOBHTTP) Name() string { return "OOBHTTP" }
func (o *OOBHTTP) Timestamp() time.Time { return time.Now() }
