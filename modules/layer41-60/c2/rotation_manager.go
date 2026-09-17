package c2

import (
	"time"
)

type RotationManager struct{}

func NewRotationManager() *RotationManager {
	return &RotationManager{}
}

func (r *RotationManager) Rotate() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "rotation_manager:done")
	return results, nil
}

func (r *RotationManager) Name() string { return "RotationManager" }
func (r *RotationManager) Timestamp() time.Time { return time.Now() }
