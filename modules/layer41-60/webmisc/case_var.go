package webmisc

import (
	"time"
)

type CaseVar struct{}

func NewCaseVar() *CaseVar {
	return &CaseVar{}
}

func (c *CaseVar) Vary() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "case_var:done")
	return results, nil
}

func (c *CaseVar) Name() string { return "CaseVar" }
func (c *CaseVar) Timestamp() time.Time { return time.Now() }
