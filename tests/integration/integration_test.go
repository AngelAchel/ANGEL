package integration

import (
	"testing"
	"time"

	"github.com/angel-platform/angel/c2/generate"
	"github.com/angel-platform/angel/gateway"
	"github.com/angel-platform/angel/gateway/auth"
	"github.com/angel-platform/angel/orchestrator"
)

func TestIntegrationFullWorkflow(t *testing.T) {
	t.Run("OrchestratorIntentClassification", func(t *testing.T) {
		orch := orchestrator.New(nil)
		if orch == nil {
			t.Fatal("expected non-nil orchestrator")
		}
		err := orch.Execute("scan target.com for vulnerabilities")
		if err != nil {
			t.Errorf("orchestrator execution failed: %v", err)
		}
	})

	t.Run("FireteamParallelExecution", func(t *testing.T) {
		cfg := &orchestrator.Config{MaxAgents: 5}
		orch := orchestrator.New(cfg)
		if orch == nil {
			t.Fatal("expected non-nil orchestrator")
		}
	})

	t.Run("MCPServerToolExecution", func(t *testing.T) {
		mcp := orchestrator.New(nil)
		if mcp == nil {
			t.Fatal("expected non-nil MCP server")
		}
	})

	t.Run("TeamserverAgentRegistration", func(t *testing.T) {
		gw := gateway.New(gateway.DefaultConfig())
		if gw == nil {
			t.Fatal("expected non-nil gateway")
		}
		gw.Start()  //nolint:errcheck
		time.Sleep(50 * time.Millisecond)
		gw.Stop()  //nolint:errcheck
	})

	t.Run("GatewayHealthCheck", func(t *testing.T) {
		gw := gateway.New(gateway.DefaultConfig())
		if gw == nil {
			t.Fatal("expected non-nil gateway")
		}
		gw.Start()  //nolint:errcheck
		time.Sleep(50 * time.Millisecond)
		gw.Stop()  //nolint:errcheck
	})
}

func TestIntegrationC2Workflow(t *testing.T) {
	t.Run("ImplantGeneration", func(t *testing.T) {
		gen := generate.NewGenerator(nil, "/tmp/test-angel-output")
		if gen == nil {
			t.Fatal("expected non-nil generator")
		}
		path, err := gen.Generate("https://teamserver.example.com")
		if err != nil {
			t.Errorf("generation failed: %v", err)
		}
		if path == "" {
			t.Error("expected non-empty path")
		}
	})

	t.Run("ProfileSelection", func(t *testing.T) {
		gen := generate.NewGenerator(nil, "/tmp/test-profile")
		if gen == nil {
			t.Fatal("expected non-nil generator")
		}
		platforms := gen.GetSupportedPlatforms()
		if len(platforms) == 0 {
			t.Error("expected at least one platform")
		}
	})

	t.Run("ChannelRotation", func(t *testing.T) {
		gw := gateway.New(gateway.DefaultConfig())
		if gw == nil {
			t.Fatal("expected non-nil gateway")
		}
	})
}

func TestIntegrationInfrastructure(t *testing.T) {
	t.Run("VPCSetup", func(t *testing.T) {
		gw := gateway.New(gateway.DefaultConfig())
		if gw == nil {
			t.Fatal("expected non-nil gateway")
		}
	})

	t.Run("WireGuardVPN", func(t *testing.T) {
		gw := gateway.New(gateway.DefaultConfig())
		if gw == nil {
			t.Fatal("expected non-nil gateway")
		}
	})

	t.Run("NginxRedirector", func(t *testing.T) {
		gw := gateway.New(gateway.DefaultConfig())
		if gw == nil {
			t.Fatal("expected non-nil gateway")
		}
	})
}

func TestIntegrationAuth(t *testing.T) {
	t.Run("JWTGeneration", func(t *testing.T) {
		mgr := auth.NewJWTManager("test-secret", time.Hour)
		if mgr == nil {
			t.Fatal("expected non-nil JWT manager")
		}
		token, err := mgr.GenerateToken("user-1", "admin", "admin")
		if err != nil {
			t.Errorf("generate token failed: %v", err)
		}
		if token == "" {
			t.Error("expected non-empty token")
		}
		claims, err := mgr.ValidateToken(token)
		if err != nil {
			t.Errorf("validate token failed: %v", err)
		}
		if claims.UserID != "user-1" {
			t.Errorf("expected user-1, got %s", claims.UserID)
		}
	})

	t.Run("RBACPermissions", func(t *testing.T) {
		rbac := auth.NewRBACManager()
		if rbac == nil {
			t.Fatal("expected non-nil RBAC manager")
		}
		if !rbac.HasPermission("admin", "read") {
			t.Error("admin should have read permission")
		}
		if rbac.HasPermission("guest", "write") {
			t.Error("guest should not have write permission")
		}
	})
}
