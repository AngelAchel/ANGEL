package webmisc

import (
	"testing"
)

func TestCommentInsert(t *testing.T) {
	c := NewComment()
	if c.Name() != "Comment" {
		t.Errorf("expected Comment, got %s", c.Name())
	}
}
