package evasion

import (
	"time"
)

type FSAccess struct{}

func NewFSAccess() *FSAccess {
	return &FSAccess{}
}

func (e *FSAccess) Access(path string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "fs_access:done")
	return results, nil
}

func (e *FSAccess) Name() string { return "FSAccess" }
func (e *FSAccess) Timestamp() time.Time { return time.Now() }
