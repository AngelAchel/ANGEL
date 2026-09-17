package evasion

import (
	"time"
)

type AggregationExfil struct{}

func NewAggregationExfil() *AggregationExfil {
	return &AggregationExfil{}
}

func (e *AggregationExfil) Exfiltrate(data string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "aggregation_exfil:done")
	return results, nil
}

func (e *AggregationExfil) Name() string { return "AggregationExfil" }
func (e *AggregationExfil) Category() EvasionCategory { return CategoryNetEvasion }
func (e *AggregationExfil) Timestamp() time.Time { return time.Now() }
