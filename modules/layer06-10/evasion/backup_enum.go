package evasion

import (
	"time"
)

type BackupEnum struct{}

func NewBackupEnum() *BackupEnum {
	return &BackupEnum{}
}

func (e *BackupEnum) Enumerate() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "backup_enum:done")
	return results, nil
}

func (e *BackupEnum) Name() string              { return "BackupEnum" }
func (e *BackupEnum) Category() EvasionCategory { return CategorySyscall }
func (e *BackupEnum) Timestamp() time.Time      { return time.Now() }
