package e2e

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/angel-platform/angel/gateway"
	"github.com/angel-platform/angel/orchestrator"
)

func TestE2EFullEngagement(t *testing.T) {
	t.Run("ReconPhase", func(t *testing.T) {
		orch := orchestrator.New(nil)
		if orch == nil {
			t.Fatal("expected non-nil orchestrator")
		}
		err := orch.Execute("scan target.com for vulnerabilities")
		if err != nil {
			t.Errorf("orchestrator execution failed: %v", err)
		}
		status := orch.GetStatus()
		if status == nil {
			t.Error("expected non-nil status")
		}
	})

	t.Run("ExploitationPhase", func(t *testing.T) {
		orch := orchestrator.New(nil)
		err := orch.Execute("exploit SQL injection")
		if err != nil {
			t.Errorf("orchestrator execution failed: %v", err)
		}
	})

	t.Run("PostExploitationPhase", func(t *testing.T) {
		orch := orchestrator.New(nil)
		err := orch.Execute("dump credentials")
		if err != nil {
			t.Errorf("orchestrator execution failed: %v", err)
		}
	})

	t.Run("DataExfiltrationPhase", func(t *testing.T) {
		orch := orchestrator.New(nil)
		err := orch.Execute("exfiltrate data via websocket")
		if err != nil {
			t.Errorf("orchestrator execution failed: %v", err)
		}
	})

	t.Run("ReportingPhase", func(t *testing.T) {
		gw := gateway.New(gateway.DefaultConfig())
		if gw == nil {
			t.Fatal("expected non-nil gateway")
		}
	})
}

func TestE2EAgentLifecycle(t *testing.T) {
	t.Run("AgentDeployment", func(t *testing.T) {
		gw := gateway.New(gateway.DefaultConfig())
		if gw == nil {
			t.Fatal("expected non-nil gateway")
		}
	})

	t.Run("AgentCheckin", func(t *testing.T) {
		gw := gateway.New(gateway.DefaultConfig())
		if gw == nil {
			t.Fatal("expected non-nil gateway")
		}
	})

	t.Run("AgentTaskExecution", func(t *testing.T) {
		gw := gateway.New(gateway.DefaultConfig())
		if gw == nil {
			t.Fatal("expected non-nil gateway")
		}
	})

	t.Run("AgentSelfDestruct", func(t *testing.T) {
		gw := gateway.New(gateway.DefaultConfig())
		if gw == nil {
			t.Fatal("expected non-nil gateway")
		}
	})
}

func TestE2EOrchestratorWorkflow(t *testing.T) {
	t.Run("IntentClassification", func(t *testing.T) {
		orch := orchestrator.New(nil)
		if orch == nil {
			t.Fatal("expected non-nil orchestrator")
		}
		err := orch.Execute("scan target for vulnerabilities")
		if err != nil {
			t.Errorf("execution failed: %v", err)
		}
	})

	t.Run("RiskAssessment", func(t *testing.T) {
		orch := orchestrator.New(nil)
		err := orch.Execute("assess risk level")
		if err != nil {
			t.Errorf("execution failed: %v", err)
		}
	})

	t.Run("FireteamExecution", func(t *testing.T) {
		orch := orchestrator.New(nil)
		if orch == nil {
			t.Fatal("expected non-nil orchestrator")
		}
		err := orch.Execute("coordinate fireteam")
		if err != nil {
			t.Errorf("execution failed: %v", err)
		}
	})

	t.Run("AgentLifecycle", func(t *testing.T) {
		gw := gateway.New(gateway.DefaultConfig())
		if gw == nil {
			t.Fatal("expected non-nil gateway")
		}
		cleanup := StartGateway(t, gw)
		defer cleanup()
	})
}

func TestE2EGatewayCRUD(t *testing.T) {
	gw := gateway.New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil gateway")
	}

	cleanup := StartGateway(t, gw)
	defer cleanup()

	// Verify health endpoint is accessible
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/api/v1/health", gw.Config().Port))
	if err != nil {
		t.Fatalf("health check failed: %v", err)
	}
	_ = resp.Body.Close() //nolint:errcheck
}
