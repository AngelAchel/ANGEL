package gateway
//nolint:staticcheck

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/angel-platform/angel/gateway/auth"
)

type Middleware interface {
	Wrap(next http.Handler) http.Handler
}

type Agent struct {
	ID        string    `json:"id"`
	Hostname  string    `json:"hostname"`
	IP        string    `json:"ip"`
	OS        string    `json:"os"`
	Arch      string    `json:"arch"`
	User      string    `json:"user"`
	Status    string    `json:"status"`
	LastCheck time.Time `json:"last_check"`
	FirstSeen time.Time `json:"first_seen"`
}

type Task struct {
	ID        string    `json:"id"`
	AgentID   string    `json:"agent_id"`
	Type      string    `json:"type"`
	Payload   string    `json:"payload"`
	Status    string    `json:"status"`
	Priority  int       `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
}

type Result struct {
	ID        string    `json:"id"`
	TaskID    string    `json:"task_id"`
	AgentID   string    `json:"agent_id"`
	Output    string    `json:"output"`
	Success   bool      `json:"success"`
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
}

type Gateway struct {
	config     *Config
	server     *http.Server
	middleware []Middleware
	routes     map[string]http.Handler
	methods    map[string]map[string]http.Handler
	mu         sync.RWMutex
	startTime  time.Time
	totalReqs  uint64
	totalErrs  uint64
	jwtMgr     *auth.JWTManager
	agents     map[string]*Agent
	tasks      map[string]*Task
	results    map[string]*Result
	activity   []string
}

type Config struct {
	Addr           string        `json:"addr"`
	Port           int           `json:"port"`
	ReadTimeout    time.Duration `json:"read_timeout"`
	WriteTimeout   time.Duration `json:"write_timeout"`
	MaxHeaderBytes int           `json:"max_header_bytes"`
	EnableTLS      bool          `json:"enable_tls"`
	TLSCertFile    string        `json:"tls_cert_file"`
	TLSKeyFile     string        `json:"tls_key_file"`
	RateLimit      int           `json:"rate_limit"`
	MaxConns       int           `json:"max_conns"`
	CORSOrigin     string        `json:"cors_origin"`
	JWTSecret      string        `json:"jwt_secret"`
	JWTExpiry      time.Duration `json:"jwt_expiry"`
	EnableRBAC     bool          `json:"enable_rbac"`
	EnableMetrics  bool          `json:"enable_metrics"`
}

func DefaultConfig() *Config {
	return &Config{
		Addr:           "0.0.0.0",
		Port:           3000,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
		RateLimit:      100,
		MaxConns:       1000,
		JWTExpiry:      24 * time.Hour,
		EnableRBAC:     true,
		EnableMetrics:  true,
	}
}

type loggingMiddleware struct{}
type recoveryMiddleware struct{}
type corsMiddleware struct{ Origin string }
type rateLimitMiddleware struct{ Rate int }
type requestIDMiddleware struct{}
type metricsMiddleware struct{}

func (m *loggingMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func (m *recoveryMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %v", err)
				http.Error(w, `{"error":"internal_server_error"}`, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (m *corsMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := m.Origin
		if origin == "" {
			origin = "*"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-ID")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *rateLimitMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", m.Rate))
		next.ServeHTTP(w, r)
	})
}

func (m *requestIDMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-ID", time.Now().Format("20060102150405.000000000"))
		next.ServeHTTP(w, r)
	})
}

func (m *metricsMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func New(cfg *Config) *Gateway {
	if cfg == nil {
		cfg = DefaultConfig()
	}
	gw := &Gateway{
		config:    cfg,
		routes:    make(map[string]http.Handler),
		methods:   make(map[string]map[string]http.Handler),
		startTime: time.Now(),
		jwtMgr:    auth.NewJWTManager(cfg.JWTSecret, cfg.JWTExpiry),
		agents:    make(map[string]*Agent),
		tasks:     make(map[string]*Task),
		results:   make(map[string]*Result),
		activity:  make([]string, 0),
	}
	gw.middleware = []Middleware{
		&loggingMiddleware{},
		&recoveryMiddleware{},
		&corsMiddleware{Origin: cfg.CORSOrigin},
		&rateLimitMiddleware{Rate: cfg.RateLimit},
		&requestIDMiddleware{},
		&metricsMiddleware{},
	}
	gw.setupRoutes()
	mux := http.NewServeMux()
	for path, handler := range gw.routes {
		mux.Handle(path, handler)
	}
	gw.server = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port),
		Handler:        gw.chainMiddleware(mux),
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
	}
	return gw
}

func (gw *Gateway) setupRoutes() {
	gw.routes["/api/v1/health"] = http.HandlerFunc(gw.handleHealth)
	gw.routes["/api/v1/auth/login"] = http.HandlerFunc(gw.handleLogin)
	gw.routes["/api/v1/auth/logout"] = http.HandlerFunc(gw.handleLogout)
	gw.routes["/api/v1/auth/refresh"] = http.HandlerFunc(gw.handleRefresh)
	gw.routes["/api/v1/agents"] = http.HandlerFunc(gw.handleAgents)
	gw.routes["/api/v1/agents/{id}"] = http.HandlerFunc(gw.handleAgentByID)
	gw.routes["/api/v1/tasks"] = http.HandlerFunc(gw.handleTasks)
	gw.routes["/api/v1/tasks/{id}"] = http.HandlerFunc(gw.handleTaskByID)
	gw.routes["/api/v1/results"] = http.HandlerFunc(gw.handleResults)
	gw.routes["/api/v1/reports"] = http.HandlerFunc(gw.handleReports)
	gw.routes["/api/v1/stats"] = http.HandlerFunc(gw.handleStats)
	gw.routes["/api/v1/activity"] = http.HandlerFunc(gw.handleActivity)
	gw.routes["/ws"] = http.HandlerFunc(gw.handleWebSocket)
}

func (gw *Gateway) chainMiddleware(next http.Handler) http.Handler {
	for i := len(gw.middleware) - 1; i >= 0; i-- {
		next = gw.middleware[i].Wrap(next)
	}
	return next
}

func (gw *Gateway) Start() error {
	log.Printf("API Gateway starting on %s:%d", gw.config.Addr, gw.config.Port)
	return gw.server.ListenAndServe()
}

func (gw *Gateway) Stop() error {
	return gw.server.Close()
}

func (gw *Gateway) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func (gw *Gateway) addActivity(action, detail string) {
	gw.mu.Lock()
	defer gw.mu.Unlock()
	gw.activity = append(gw.activity, fmt.Sprintf("[%s] %s: %s", time.Now().Format("2006-01-02 15:04:05"), action, detail))
	if len(gw.activity) > 100 {
		gw.activity = gw.activity[1:]
	}
}

func (gw *Gateway) handleHealth(w http.ResponseWriter, r *http.Request) {
	gw.writeJSON(w, http.StatusOK, map[string]interface{}{
		"status": "healthy",
		"uptime": time.Since(gw.startTime).String(),
		"agents": len(gw.agents),
		"tasks":  len(gw.tasks),
	})
}

func (gw *Gateway) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		gw.writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method_not_allowed"})
		return
	}
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		gw.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid_request"})
		return
	}
	if creds.Username == "" || creds.Password == "" {
		gw.writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing_credentials"})
		return
	}
	role := "viewer"
	if creds.Username == "admin" {
		role = "admin"
	} else if creds.Username == "operator" {
		role = "operator"
	}
	token, err := gw.jwtMgr.GenerateToken("user", creds.Username, role)
	if err != nil {
		gw.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "token_generation_failed"})
		return
	}
	gw.addActivity("auth", fmt.Sprintf("user %s logged in as %s", creds.Username, role))
	gw.writeJSON(w, http.StatusOK, map[string]interface{}{
		"token":      token,
		"expires_in": int(gw.config.JWTExpiry.Seconds()),
	})
}

func (gw *Gateway) handleLogout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		gw.writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method_not_allowed"})
		return
	}
	gw.addActivity("auth", "user logged out")
	gw.writeJSON(w, http.StatusOK, map[string]string{"status": "logged_out"})
}

func (gw *Gateway) handleRefresh(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		gw.writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method_not_allowed"})
		return
	}
	token, err := gw.jwtMgr.GenerateToken("user", "admin", "operator")
	if err != nil {
		gw.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "token_generation_failed"})
		return
	}
	gw.writeJSON(w, http.StatusOK, map[string]interface{}{
		"token":      token,
		"expires_in": int(gw.config.JWTExpiry.Seconds()),
	})
}

func (gw *Gateway) handleAgents(w http.ResponseWriter, r *http.Request) {
	gw.mu.RLock()
	agentList := make([]*Agent, 0, len(gw.agents))
	for _, a := range gw.agents {
		agentList = append(agentList, a)
	}
	gw.mu.RUnlock()
	switch r.Method {
	case http.MethodGet:
		gw.writeJSON(w, http.StatusOK, map[string]interface{}{"agents": agentList})
	case http.MethodPost:
		var req struct {
			Hostname string `json:"hostname"`
			IP       string `json:"ip"`
			OS       string `json:"os"`
			Arch     string `json:"arch"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			gw.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid_request"})
			return
		}
		agent := &Agent{
			ID:        fmt.Sprintf("agent-%d", time.Now().UnixNano()),
			Hostname:  req.Hostname,
			IP:        req.IP,
			OS:        req.OS,
			Arch:      req.Arch,
			Status:    "online",
			LastCheck: time.Now(),
			FirstSeen: time.Now(),
		}
		gw.mu.Lock()
		gw.agents[agent.ID] = agent
		gw.mu.Unlock()
		gw.addActivity("agents", fmt.Sprintf("registered %s (%s)", req.Hostname, req.IP))
		gw.writeJSON(w, http.StatusCreated, map[string]interface{}{"status": "agent_registered", "agent_id": agent.ID})
	case http.MethodDelete:
		gw.mu.Lock()
		gw.agents = make(map[string]*Agent)
		gw.mu.Unlock()
		gw.addActivity("agents", "all agents killed")
		gw.writeJSON(w, http.StatusOK, map[string]string{"status": "all_agents_killed"})
	default:
		gw.writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method_not_allowed"})
	}
}

func (gw *Gateway) handleAgentByID(w http.ResponseWriter, r *http.Request) {
	gw.mu.RLock()
	var agent *Agent
	for _, a := range gw.agents {
		agent = a
		break
	}
	gw.mu.RUnlock()
	if agent == nil {
		gw.writeJSON(w, http.StatusNotFound, map[string]string{"error": "agent_not_found"})
		return
	}
	gw.writeJSON(w, http.StatusOK, agent)
}

func (gw *Gateway) handleTasks(w http.ResponseWriter, r *http.Request) {
	gw.mu.RLock()
	taskList := make([]*Task, 0, len(gw.tasks))
	for _, t := range gw.tasks {
		taskList = append(taskList, t)
	}
	gw.mu.RUnlock()
	switch r.Method {
	case http.MethodGet:
		gw.writeJSON(w, http.StatusOK, map[string]interface{}{"tasks": taskList})
	case http.MethodPost:
		var req struct {
			AgentID string `json:"agent_id"`
			Command string `json:"command"`
			Args    string `json:"arguments"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			gw.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid_request"})
			return
		}
		task := &Task{
			ID:        fmt.Sprintf("task-%d", time.Now().UnixNano()),
			AgentID:   req.AgentID,
			Type:      "command",
			Payload:   fmt.Sprintf("%s %s", req.Command, req.Args),
			Status:    "pending",
			Priority:  1,
			CreatedAt: time.Now(),
		}
		gw.mu.Lock()
		gw.tasks[task.ID] = task
		gw.mu.Unlock()
		gw.addActivity("tasks", fmt.Sprintf("created task %s for agent %s", task.ID, req.AgentID))
		gw.writeJSON(w, http.StatusCreated, map[string]interface{}{"status": "task_created", "task_id": task.ID})
	default:
		gw.writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method_not_allowed"})
	}
}

func (gw *Gateway) handleTaskByID(w http.ResponseWriter, r *http.Request) {
	gw.mu.RLock()
	var task *Task
	for _, t := range gw.tasks {
		task = t
		break
	}
	gw.mu.RUnlock()
	if task == nil {
		gw.writeJSON(w, http.StatusNotFound, map[string]string{"error": "task_not_found"})
		return
	}
	gw.writeJSON(w, http.StatusOK, task)
}

func (gw *Gateway) handleResults(w http.ResponseWriter, r *http.Request) {
	gw.mu.RLock()
	resultList := make([]*Result, 0, len(gw.results))
	for _, r := range gw.results {
		resultList = append(resultList, r)
	}
	gw.mu.RUnlock()
	gw.writeJSON(w, http.StatusOK, map[string]interface{}{"results": resultList})
}

func (gw *Gateway) handleReports(w http.ResponseWriter, r *http.Request) {
	gw.mu.RLock()
	reportCount := len(gw.results)
	gw.mu.RUnlock()
	gw.writeJSON(w, http.StatusOK, map[string]interface{}{
		"reports": []map[string]interface{}{
			{"id": "report-1", "type": "executive", "status": "completed"},
			{"id": "report-2", "type": "technical", "status": "pending"},
		},
		"total": reportCount,
	})
}

func (gw *Gateway) handleStats(w http.ResponseWriter, r *http.Request) {
	gw.mu.RLock()
	activeAgents := 0
	pendingTasks := 0
	completedTasks := 0
	for _, a := range gw.agents {
		if a.Status == "online" {
			activeAgents++
		}
	}
	for _, t := range gw.tasks {
		if t.Status == "pending" {
			pendingTasks++
		} else if t.Status == "done" {
			completedTasks++
		}
	}
	gw.mu.RUnlock()
	gw.writeJSON(w, http.StatusOK, map[string]interface{}{
		"active_agents":     activeAgents,
		"pending_tasks":     pendingTasks,
		"completed_tasks":   completedTasks,
		"credentials_found": 0,
	})
}

func (gw *Gateway) handleActivity(w http.ResponseWriter, r *http.Request) {
	gw.mu.RLock()
	activity := make([]string, len(gw.activity))
	copy(activity, gw.activity)
	gw.mu.RUnlock()
	gw.writeJSON(w, http.StatusOK, map[string]interface{}{"activities": activity})
}

func (gw *Gateway) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	gw.writeJSON(w, http.StatusOK, map[string]interface{}{
		"status": "websocket_endpoint",
		"url":    "ws://localhost:3000/ws",
	})
}

func (gw *Gateway) IncrementRequests() {
	gw.mu.Lock()
	defer gw.mu.Unlock()
	gw.totalReqs++
}
