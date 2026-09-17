package nosql

import (
	"time"
)

type CQLInject struct{}

func NewCQLInject() *CQLInject {
	return &CQLInject{}
}

func (c *CQLInject) Inject() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "cql_inject:done")
	return results, nil
}

func (c *CQLInject) Name() string         { return "CQLInject" }
func (c *CQLInject) Timestamp() time.Time { return time.Now() }
