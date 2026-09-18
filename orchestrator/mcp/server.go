package mcp

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type MCPServer struct {
	tools    map[string]*Tool
	sessions map[string]*Session
	mu       sync.RWMutex
	maxSess  int
}

type Tool struct {
	ID          string
	Name        string
	Description string
	Category    string
	Execute     func(params map[string]interface{}) (*ToolResult, error)
	Timeout     time.Duration
}

type ToolResult struct {
	Success  bool
	Data     map[string]interface{}
	Error    string
	Duration time.Duration
}

type Session struct {
	ID        string
	Tools     []string
	CreatedAt time.Time
	LastUsed  time.Time
	Status    string
}

func NewMCPServer(maxSessions int) *MCPServer {
	srv := &MCPServer{
		tools:    make(map[string]*Tool),
		sessions: make(map[string]*Session),
		maxSess:  maxSessions,
	}
	srv.registerDefaultTools()
	return srv
}

func (s *MCPServer) registerDefaultTools() {
	defaultTools := []*Tool{
		{ID: "nmap", Name: "Nmap Scanner", Description: "Network port scanner", Category: "recon", Timeout: 60 * time.Second,
			Execute: s.execCommand("nmap")},
		{ID: "nuclei", Name: "Nuclei Scanner", Description: "Vulnerability scanner", Category: "exploit", Timeout: 120 * time.Second,
			Execute: s.execCommand("nuclei")},
		{ID: "ffuf", Name: "FFuf Fuzzer", Description: "Web fuzzer", Category: "recon", Timeout: 60 * time.Second,
			Execute: s.execCommand("ffuf")},
		{ID: "hydra", Name: "Hydra Brute Force", Description: "Password brute force", Category: "credential", Timeout: 300 * time.Second,
			Execute: s.execCommand("hydra")},
		{ID: "metasploit", Name: "Metasploit Framework", Description: "Exploitation framework", Category: "exploit", Timeout: 600 * time.Second,
			Execute: s.execCommand("msfconsole")},
		{ID: "playwright", Name: "Playwright Browser", Description: "Browser automation", Category: "recon", Timeout: 60 * time.Second,
			Execute: s.execCommand("playwright")},
		{ID: "subfinder", Name: "Subfinder", Description: "Subdomain enumeration", Category: "recon", Timeout: 60 * time.Second,
			Execute: s.execCommand("subfinder")},
		{ID: "shodan", Name: "Shodan CLI", Description: "Internet device search", Category: "recon", Timeout: 30 * time.Second,
			Execute: s.execCommand("shodan")},
		{ID: "sqlmap", Name: "SQLMap", Description: "SQL injection tool", Category: "exploit", Timeout: 120 * time.Second,
			Execute: s.execCommand("sqlmap")},
		{ID: "mimikatz", Name: "Mimikatz", Description: "Credential extraction", Category: "credential", Timeout: 60 * time.Second,
			Execute: s.execCommand("mimikatz")},
	}

	for _, tool := range defaultTools {
		s.tools[tool.ID] = tool
	}
}

// execCommand returns an Execute function that invokes the external binary
// identified by binary. argv/flags are passed through params["args"] (string slice).
// If the binary is not found on PATH, a clear error is returned instead of the
// previous "simulated" stub result.
func (s *MCPServer) execCommand(binary string) func(params map[string]interface{}) (*ToolResult, error) {
	return func(params map[string]interface{}) (*ToolResult, error) {
		// build argv
		var argv []string
		argv = append(argv, binary)
		if raw, ok := params["args"]; ok {
			if arr, ok := raw.([]interface{}); ok {
				for _, a := range arr {
					argv = append(argv, fmt.Sprintf("%v", a))
				}
			} else if sArr, ok := raw.(string); ok {
				argv = append(argv, sArr)
			}
		}
		return s.runExternal(argv, params)
	}
}

// runExternal executes the given argv (binary + args) with the tool's timeout.
// Returns a structured ToolResult. If the binary is missing from PATH, returns
// a clear actionable error rather than the old "simulated" stub output.
func (s *MCPServer) runExternal(argv []string, params map[string]interface{}) (*ToolResult, error) {
	binary := argv[0]
	start := time.Now()

	// locate binary on PATH
	if _, err := exec.LookPath(binary); err != nil {
		return &ToolResult{
			Success: false,
			Error:   fmt.Sprintf("binary %q not found on PATH; install it or add it to PATH", binary),
		}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, argv[0], argv[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	// optional working dir / env passthrough
	if wd, ok := params["cwd"].(string); ok && wd != "" {
		cmd.Dir = wd
	}
	if envMap, ok := params["env"].(map[string]interface{}); ok {
		cmd.Env = append(cmd.Environ())
		for k, v := range envMap {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%v", k, v))
		}
	}

	err := cmd.Run()
	result := &ToolResult{
		Success: err == nil,
		Data: map[string]interface{}{
			"tool":   binary,
			"argv":   strings.Join(argv, " "),
			"stdout": stdout.String(),
			"stderr": stderr.String(),
		},
		Duration: time.Since(start),
	}
	if err != nil {
		result.Error = fmt.Sprintf("exit=%v: %s", err, stderr.String())
		// context.DeadlineExceeded → keep success=false with a timeout note
	}
	if ctx.Err() == context.DeadlineExceeded {
		result.Error = fmt.Sprintf("timeout after %v", 120*time.Second)
		result.Success = false
	}
	return result, nil
}

func (s *MCPServer) RegisterTool(tool *Tool) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tools[tool.ID]; exists {
		return fmt.Errorf("tool %s already registered", tool.ID)
	}

	s.tools[tool.ID] = tool
	return nil
}

func (s *MCPServer) ExecuteTool(toolID string, params map[string]interface{}) (*ToolResult, error) {
	s.mu.RLock()
	tool, exists := s.tools[toolID]
	s.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("tool %s not found", toolID)
	}

	start := time.Now()

	var result *ToolResult
	var err error

	if tool.Execute != nil {
		result, err = tool.Execute(params)
		if err != nil {
			return &ToolResult{
				Success:  false,
				Error:    err.Error(),
				Duration: time.Since(start),
			}, nil
		}
	} else {
		result = &ToolResult{
			Success: true,
			Data: map[string]interface{}{
				"tool":    toolID,
				"status":  "simulated",
				"message": fmt.Sprintf("Tool %s executed with params: %v", toolID, params),
			},
			Duration: time.Since(start),
		}
	}

	return result, nil
}

func (s *MCPServer) CreateSession(sessionID string) (*Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.sessions) >= s.maxSess {
		return nil, fmt.Errorf("max sessions reached (%d)", s.maxSess)
	}

	session := &Session{
		ID:        sessionID,
		Tools:     make([]string, 0),
		CreatedAt: time.Now(),
		LastUsed:  time.Now(),
		Status:    "active",
	}

	s.sessions[sessionID] = session
	return session, nil
}

func (s *MCPServer) CloseSession(sessionID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	session, exists := s.sessions[sessionID]
	if !exists {
		return fmt.Errorf("session %s not found", sessionID)
	}

	session.Status = "closed"
	delete(s.sessions, sessionID)
	return nil
}

func (s *MCPServer) ListTools() []*Tool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tools := make([]*Tool, 0, len(s.tools))
	for _, tool := range s.tools {
		tools = append(tools, tool)
	}
	return tools
}

func (s *MCPServer) GetTool(toolID string) (*Tool, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	tool, exists := s.tools[toolID]
	return tool, exists
}

func (s *MCPServer) GetStatus() map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return map[string]interface{}{
		"total_tools":     len(s.tools),
		"active_sessions": len(s.sessions),
		"max_sessions":    s.maxSess,
	}
}
