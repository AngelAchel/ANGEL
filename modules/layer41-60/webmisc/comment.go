package webmisc

import (
	"time"
)

type Comment struct{}

func NewComment() *Comment {
	return &Comment{}
}

func (c *Comment) Insert() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "comment:done")
	return results, nil
}

func (c *Comment) Name() string         { return "Comment" }
func (c *Comment) Timestamp() time.Time { return time.Now() }
