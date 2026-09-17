package c2server

import (
	"time"
)

type Ldap struct{}

func NewLdap() *Ldap {
	return &Ldap{}
}

func (e *Ldap) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "ldap:done")
	return results, nil
}

func (e *Ldap) Name() string { return "Ldap" }
func (e *Ldap) Timestamp() time.Time { return time.Now() }
