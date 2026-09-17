package c2server

import (
	"time"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (e *Parser) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "parser:done")
	return results, nil
}

func (e *Parser) Name() string { return "Parser" }
func (e *Parser) Timestamp() time.Time { return time.Now() }
