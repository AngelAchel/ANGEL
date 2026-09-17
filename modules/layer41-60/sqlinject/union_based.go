package sqlinject

import (
	"time"
)

type UnionBased struct{}

func NewUnionBased() *UnionBased {
	return &UnionBased{}
}

func (u *UnionBased) Inject() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "union_based:done")
	return results, nil
}

func (u *UnionBased) Name() string         { return "UnionBased" }
func (u *UnionBased) Timestamp() time.Time { return time.Now() }
