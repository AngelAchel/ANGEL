package nosql

import (
	"time"
)

type LookupExfil struct{}

func NewLookupExfil() *LookupExfil {
	return &LookupExfil{}
}

func (l *LookupExfil) Exfiltrate() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "lookup_exfil:done")
	return results, nil
}

func (l *LookupExfil) Name() string { return "LookupExfil" }
func (l *LookupExfil) Timestamp() time.Time { return time.Now() }
