package evasion

import (
	"time"
)

type UDFInstall struct{}

func NewUDFInstall() *UDFInstall {
	return &UDFInstall{}
}

func (e *UDFInstall) Install(funcName string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "udf_install:done")
	return results, nil
}

func (e *UDFInstall) Name() string { return "UDFInstall" }
func (e *UDFInstall) Timestamp() time.Time { return time.Now() }
