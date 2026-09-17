package webmisc

import (
	"time"
)

type Multipart struct{}

func NewMultipart() *Multipart {
	return &Multipart{}
}

func (m *Multipart) Upload() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "multipart:done")
	return results, nil
}

func (m *Multipart) Name() string         { return "Multipart" }
func (m *Multipart) Timestamp() time.Time { return time.Now() }
