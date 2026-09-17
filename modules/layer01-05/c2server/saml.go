package c2server

import (
	"time"
)

type Saml struct{}

func NewSaml() *Saml {
	return &Saml{}
}

func (e *Saml) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "saml:done")
	return results, nil
}

func (e *Saml) Name() string { return "Saml" }
func (e *Saml) Timestamp() time.Time { return time.Now() }
