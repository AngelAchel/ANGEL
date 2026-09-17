package evasion

import (
	"time"
)

type CharFunc struct{}

func NewCharFunc() *CharFunc {
	return &CharFunc{}
}

func (e *CharFunc) Process(char string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "char_func:processed")
	return results, nil
}

func (e *CharFunc) Name() string              { return "CharFunc" }
func (e *CharFunc) Category() EvasionCategory { return CategorySyscall }
func (e *CharFunc) Timestamp() time.Time      { return time.Now() }
