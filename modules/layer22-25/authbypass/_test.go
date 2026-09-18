package authbypass

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestAuthBypassEngineTestBypass(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.TestBypass()
	if err != nil {
		t.Errorf("TestBypass failed: %v", err)
	}
}

func TestAuthBypassEngineFullBypass(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.FullBypass()
	if err != nil {
		t.Errorf("FullBypass failed: %v", err)
	}
}

func TestAuthBypassEnginetestSQLiAuth(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.testSQLiAuth()
	if err != nil {
		t.Errorf("testSQLiAuth failed: %v", err)
	}
}

func TestAuthBypassEngineRun(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	result, err := gw.Run()
	if err != nil {
		t.Errorf("Run failed: %v", err)
	}
	if result == "" {
		t.Error("expected non-empty result")
	}
}
