package webmisc

import (
	"time"
)

type Whitespace struct{}

func NewWhitespace() *Whitespace {
	return &Whitespace{}
}

func (w *Whitespace) Obfuscate() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "whitespace:done")
	return results, nil
}

func (w *Whitespace) Name() string         { return "Whitespace" }
func (w *Whitespace) Timestamp() time.Time { return time.Now() }
