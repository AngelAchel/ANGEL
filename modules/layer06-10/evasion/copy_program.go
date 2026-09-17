package evasion

import (
	"time"
)

type CopyProgram struct{}

func NewCopyProgram() *CopyProgram {
	return &CopyProgram{}
}

func (e *CopyProgram) Copy(src string, dst string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "copy_program:done")
	return results, nil
}

func (e *CopyProgram) Name() string { return "CopyProgram" }
func (e *CopyProgram) Category() EvasionCategory { return CategorySyscall }
func (e *CopyProgram) Timestamp() time.Time { return time.Now() }
