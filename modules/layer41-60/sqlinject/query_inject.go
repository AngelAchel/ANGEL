package sqlinject

import (
	"fmt"
	"strings"
)

// QueryInjectConfig holds parameters for query DSL injection (Elasticsearch/Solr).
type QueryInjectConfig struct {
	TargetURL  string   `json:"target_url"`
	QueryType  string   `json:"query_type"`
	Fields     []string `json:"fields"`
	Index      string   `json:"index"`
	Collection string   `json:"collection"`
}

// QueryResult holds the output of a query DSL injection operation.
type QueryResult struct {
	Success  bool   `json:"success"`
	Method   string `json:"method"`
	Message  string `json:"message"`
	Payload  string `json:"payload"`
	QueryDSL string `json:"query_dsl"`
}

// Engine is the query DSL injection engine for Elasticsearch/Solr.
type QueryEngine struct {
	config QueryInjectConfig
}

// NewEngine creates a new query DSL injection engine.
func NewEngine(cfg QueryInjectConfig) *QueryEngine {
	return &QueryEngine{config: cfg}
}

// ElasticsearchMatch builds an Elasticsearch match query injection payload.
func (e *QueryEngine) ElasticsearchMatch(field string, value string) (string, error) {
	if field == "" || value == "" {
		return "", fmt.Errorf("field and value are required")
	}

	dsl := fmt.Sprintf(`{"query":{"match":{"%s":"%s"}}}`, field, value)
	payload := fmt.Sprintf("%s'/%s/*", e.config.TargetURL, dsl)

	return payload, nil
}

// ElasticsearchBool builds an Elasticsearch bool query injection payload.
func (e *QueryEngine) ElasticsearchBool(mustClauses []string, filterClauses []string) (string, error) {
	if len(mustClauses) == 0 && len(filterClauses) == 0 {
		return "", fmt.Errorf("at least one clause required")
	}

	boolQuery := `{"query":{"bool":{`
	parts := make([]string, 0)
	if len(mustClauses) > 0 {
		parts = append(parts, fmt.Sprintf(`"must":[%s]`, strings.Join(mustClauses, ",")))
	}
	if len(filterClauses) > 0 {
		parts = append(parts, fmt.Sprintf(`"filter":[%s]`, strings.Join(filterClauses, ",")))
	}
	boolQuery += strings.Join(parts, ",") + "}}}}"

	payload := fmt.Sprintf("%s/%s/_search", e.config.TargetURL, e.config.Index)
	_ = boolQuery // DSL would be sent in request body

	return payload, nil
}

// ElasticsearchRange builds an Elasticsearch range query injection payload.
func (e *QueryEngine) ElasticsearchRange(field string, gte string, lte string) (string, error) {
	if field == "" {
		return "", fmt.Errorf("field is required")
	}

	dsl := fmt.Sprintf(`{"query":{"range":{"%s":{"gte":"%s","lte":"%s"}}}}`, field, gte, lte)
	payload := fmt.Sprintf("%s/%s/_search", e.config.TargetURL, e.config.Index)
	_ = dsl // DSL would be sent in request body

	return payload, nil
}

// ElasticsearchScript builds an Elasticsearch script query injection payload.
func (e *QueryEngine) ElasticsearchScript(field string, script string) (string, error) {
	if field == "" || script == "" {
		return "", fmt.Errorf("field and script are required")
	}

	dsl := fmt.Sprintf(`{"query":{"script":{"script":{"source":"%s","params":{"field":"%s"}}}}`, script, field)
	payload := fmt.Sprintf("%s/%s/_search", e.config.TargetURL, e.config.Index)
	_ = dsl // DSL would be sent in request body

	return payload, nil
}

// SolrQuery builds a Solr query injection payload.
func (e *QueryEngine) SolrQuery(field string, value string) (string, error) {
	if field == "" || value == "" {
		return "", fmt.Errorf("field and value are required")
	}

	payload := fmt.Sprintf("%s/solr/%s/select?q=%s:%s", e.config.TargetURL, e.config.Collection, field, value)
	return payload, nil
}

// SolrFuzzy builds a Solr fuzzy query injection payload.
func (e *QueryEngine) SolrFuzzy(field string, value string) (string, error) {
	if field == "" || value == "" {
		return "", fmt.Errorf("field and value are required")
	}

	payload := fmt.Sprintf("%s/solr/%s/select?q=%s:%s~", e.config.TargetURL, e.config.Collection, field, value)
	return payload, nil
}

// SolrWildcard builds a Solr wildcard query injection payload.
func (e *QueryEngine) SolrWildcard(field string, pattern string) (string, error) {
	if field == "" || pattern == "" {
		return "", fmt.Errorf("field and pattern are required")
	}

	payload := fmt.Sprintf("%s/solr/%s/select?q=%s:%s*", e.config.TargetURL, e.config.Collection, field, pattern)
	return payload, nil
}

// SolrRegex builds a Solr regex query injection payload.
func (e *QueryEngine) SolrRegex(field string, regex string) (string, error) {
	if field == "" || regex == "" {
		return "", fmt.Errorf("field and regex are required")
	}

	payload := fmt.Sprintf("%s/solr/%s/select?q=%s:/%s/", e.config.TargetURL, e.config.Collection, field, regex)
	return payload, nil
}

// SolrBoost builds a Solr boost query injection payload.
func (e *QueryEngine) SolrBoost(field string, value string, boost float64) (string, error) {
	if field == "" || value == "" {
		return "", fmt.Errorf("field and value are required")
	}

	payload := fmt.Sprintf("%s/solr/%s/select?q=%s:%s^%.1f", e.config.TargetURL, e.config.Collection, field, value, boost)
	return payload, nil
}

// SolrDisMax builds a Solr dismax query injection payload.
func (e *QueryEngine) SolrDisMQ(q string,qf string,mm string) (string, error) {
	if q == "" {
		return "", fmt.Errorf("query string is required")
	}

	payload := fmt.Sprintf("%s/solr/%s/select?q=%s&defType=dismax&qf=%s&mm=%s", e.config.TargetURL, e.config.Collection, q, qf, mm)
	return payload, nil
}

// ESFunctionScore builds an Elasticsearch function_score query injection payload.
func (e *QueryEngine) ESFunctionScore(field string, boost float64) (string, error) {
	if field == "" {
		return "", fmt.Errorf("field is required")
	}

	dsl := fmt.Sprintf(`{"query":{"function_score":{"query":{"match_all":{}},"functions":[{"field_value_factor":{"field":"%s","factor":%.1f}}]}}}`, field, boost)
	payload := fmt.Sprintf("%s/%s/_search", e.config.TargetURL, e.config.Index)
	_ = dsl // DSL would be sent in request body

	return payload, nil
}

// ESScriptScore builds an Elasticsearch script_score query injection payload.
func (e *QueryEngine) ESScriptScore(script string) (string, error) {
	if script == "" {
		return "", fmt.Errorf("script cannot be empty")
	}

	dsl := fmt.Sprintf(`{"query":{"function_score":{"query":{"match_all":{}},"functions":[{"script_score":{"script":{"source":"%s"}}}]}}}`, script)
	payload := fmt.Sprintf("%s/%s/_search", e.config.TargetURL, e.config.Index)
	_ = dsl // DSL would be sent in request body

	return payload, nil
}

// ESNested builds an Elasticsearch nested query injection payload.
func (e *QueryEngine) ESNested(path string, query string) (string, error) {
	if path == "" || query == "" {
		return "", fmt.Errorf("path and query are required")
	}

	dsl := fmt.Sprintf(`{"query":{"nested":{"path":"%s","query":{%s}}}}`, path, query)
	payload := fmt.Sprintf("%s/%s/_search", e.config.TargetURL, e.config.Index)
	_ = dsl // DSL would be sent in request body

	return payload, nil
}

// ESMultiMatch builds an Elasticsearch multi_match query injection payload.
func (e *QueryEngine) ESMultiMatch(fields []string, query string) (string, error) {
	if len(fields) == 0 || query == "" {
		return "", fmt.Errorf("fields and query are required")
	}

	dsl := fmt.Sprintf(`{"query":{"multi_match":{"query":"%s","fields":[%s]}}}`, query, strings.Join(fields, ","))
	payload := fmt.Sprintf("%s/%s/_search", e.config.TargetURL, e.config.Index)
	_ = dsl // DSL would be sent in request body

	return payload, nil
}

// ESMoreLikeThis builds an Elasticsearch more_like_this query injection payload.
func (e *QueryEngine) ESMoreLikeThis(field string, likeText string) (string, error) {
	if field == "" || likeText == "" {
		return "", fmt.Errorf("field and like_text are required")
	}

	dsl := fmt.Sprintf(`{"query":{"more_like_this":{"fields":["%s"],"like_text":"%s","min_term_freq":1,"max_query_terms":12}}}`, field, likeText)
	payload := fmt.Sprintf("%s/%s/_search", e.config.TargetURL, e.config.Index)
	_ = dsl // DSL would be sent in request body

	return payload, nil
}

// ESTerm builds an Elasticsearch term query injection payload.
func (e *QueryEngine) ESTerm(field string, value string) (string, error) {
	if field == "" || value == "" {
		return "", fmt.Errorf("field and value are required")
	}

	dsl := fmt.Sprintf(`{"query":{"term":{"%s":"%s"}}}`, field, value)
	payload := fmt.Sprintf("%s/%s/_search", e.config.TargetURL, e.config.Index)
	_ = dsl // DSL would be sent in request body

	return payload, nil
}

// SolrFacet builds a Solr facet query injection payload.
func (e *QueryEngine) SolrFacet(field string, limit int) (string, error) {
	if field == "" {
		return "", fmt.Errorf("field is required")
	}
	if limit <= 0 {
		limit = 10
	}

	payload := fmt.Sprintf("%s/solr/%s/select?q=*:*&facet=true&facet.field=%s&facet.limit=%d", e.config.TargetURL, e.config.Collection, field, limit)
	return payload, nil
}

// SolrHighlight builds a Solr highlighting query injection payload.
func (e *QueryEngine) SolrHighlight(field string, query string) (string, error) {
	if field == "" || query == "" {
		return "", fmt.Errorf("field and query are required")
	}

	payload := fmt.Sprintf("%s/solr/%s/select?q=%s&hl=true&hl.fl=%s", e.config.TargetURL, e.config.Collection, query, field)
	return payload, nil
}

// SolrStats builds a Solr stats query injection payload.
func (e *QueryEngine) SolrStats(field string) (string, error) {
	if field == "" {
		return "", fmt.Errorf("field is required")
	}

	payload := fmt.Sprintf("%s/solr/%s/select?q=*:*&stats=true&stats.field=%s", e.config.TargetURL, e.config.Collection, field)
	return payload, nil
}

// ESExists builds an Elasticsearch exists query injection payload.
func (e *QueryEngine) ESExists(field string) (string, error) {
	if field == "" {
		return "", fmt.Errorf("field is required")
	}

	dsl := fmt.Sprintf(`{"query":{"exists":{"field":"%s"}}}`, field)
	payload := fmt.Sprintf("%s/%s/_search", e.config.TargetURL, e.config.Index)
	_ = dsl // DSL would be sent in request body

	return payload, nil
}

// ESIDs builds an Elasticsearch IDs query injection payload.
func (e *QueryEngine) ESIDs(ids []string) (string, error) {
	if len(ids) == 0 {
		return "", fmt.Errorf("ids cannot be empty")
	}

	dsl := fmt.Sprintf(`{"query":{"ids":{"values":[%s]}}}`, strings.Join(ids, ","))
	payload := fmt.Sprintf("%s/%s/_search", e.config.TargetURL, e.config.Index)
	_ = dsl // DSL would be sent in request body

	return payload, nil
}