package webmisc

import (
	"time"
)

type Unicode struct{}

func NewUnicode() *Unicode {
	return &Unicode{}
}

func (u *Unicode) Normalize() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "unicode:done")
	return results, nil
}

func (u *Unicode) Name() string         { return "Unicode" }
func (u *Unicode) Timestamp() time.Time { return time.Now() }
