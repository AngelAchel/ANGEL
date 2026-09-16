package gateway

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGatewayHealth(t *testing.T) {
	gw := New(DefaultConfig())
	req := httptest.NewRequest("GET", "/api/v1/health", nil)
	w := httptest.NewRecorder()
	gw.handleHealth(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGatewayAgents(t *testing.T) {
	gw := New(DefaultConfig())
	req := httptest.NewRequest("GET", "/api/v1/agents", nil)
	w := httptest.NewRecorder()
	gw.handleAgents(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGatewayAgentsPost(t *testing.T) {
	gw := New(DefaultConfig())
	body := map[string]string{"hostname": "test-host", "ip": "192.168.1.1", "os": "linux", "arch": "amd64"}
	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/api/v1/agents", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	gw.handleAgents(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}
}

func TestGatewayTasks(t *testing.T) {
	gw := New(DefaultConfig())
	req := httptest.NewRequest("GET", "/api/v1/tasks", nil)
	w := httptest.NewRecorder()
	gw.handleTasks(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGatewayTasksPost(t *testing.T) {
	gw := New(DefaultConfig())
	body := map[string]string{"agent_id": "agent-1", "command": "shell", "arguments": "ls"}
	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/api/v1/tasks", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	gw.handleTasks(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}
}

func TestGatewayStats(t *testing.T) {
	gw := New(DefaultConfig())
	req := httptest.NewRequest("GET", "/api/v1/stats", nil)
	w := httptest.NewRecorder()
	gw.handleStats(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGatewayLogin(t *testing.T) {
	gw := New(DefaultConfig())
	body := map[string]string{"username": "admin", "password": "test123"}
	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	gw.handleLogin(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGatewayLoginMissingCreds(t *testing.T) {
	gw := New(DefaultConfig())
	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewReader([]byte("{}")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	gw.handleLogin(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", w.Code)
	}
}

func TestGatewayLogout(t *testing.T) {
	gw := New(DefaultConfig())
	req := httptest.NewRequest("POST", "/api/v1/auth/logout", nil)
	w := httptest.NewRecorder()
	gw.handleLogout(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGatewayResults(t *testing.T) {
	gw := New(DefaultConfig())
	req := httptest.NewRequest("GET", "/api/v1/results", nil)
	w := httptest.NewRecorder()
	gw.handleResults(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGatewayReports(t *testing.T) {
	gw := New(DefaultConfig())
	req := httptest.NewRequest("GET", "/api/v1/reports", nil)
	w := httptest.NewRecorder()
	gw.handleReports(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGatewayActivity(t *testing.T) {
	gw := New(DefaultConfig())
	req := httptest.NewRequest("GET", "/api/v1/activity", nil)
	w := httptest.NewRecorder()
	gw.handleActivity(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGatewayMethodNotAllowed(t *testing.T) {
	gw := New(DefaultConfig())
	req := httptest.NewRequest("PATCH", "/api/v1/agents", nil)
	w := httptest.NewRecorder()
	gw.handleAgents(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", w.Code)
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.Addr != "0.0.0.0" {
		t.Errorf("expected 0.0.0.0, got %s", cfg.Addr)
	}
	if cfg.Port != 3000 {
		t.Errorf("expected 3000, got %d", cfg.Port)
	}
}
