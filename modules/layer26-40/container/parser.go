package container

import (
	"time"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "parser:done")
	return results, nil
}

func (p *Parser) Name() string         { return "Parser" }
func (p *Parser) Timestamp() time.Time { return time.Now() }
