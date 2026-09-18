package lateral

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestLateralEngineregisterMethods(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.registerMethods()
	if err != nil {
		t.Errorf("registerMethods failed: %v", err)
	}
}

func TestLateralEngineExecute(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Execute()
	if err != nil {
		t.Errorf("Execute failed: %v", err)
	}
}

func TestLateralEngineExecuteWithFallback(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ExecuteWithFallback()
	if err != nil {
		t.Errorf("ExecuteWithFallback failed: %v", err)
	}
}

func TestLateralEngineRun(t *testing.T) {
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
