package webmisc

import (
	"time"
)

type XMLParam struct{}

func NewXMLParam() *XMLParam {
	return &XMLParam{}
}

func (x *XMLParam) Inject() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "xml_param:done")
	return results, nil
}

func (x *XMLParam) Name() string         { return "XMLParam" }
func (x *XMLParam) Timestamp() time.Time { return time.Now() }
