package destruction

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestDestructionEngineExecuteChain(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ExecuteChain()
	if err != nil {
		t.Errorf("ExecuteChain failed: %v", err)
	}
}

func TestDestructionEngineCalculateBlastRadius(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.CalculateBlastRadius()
	if err != nil {
		t.Errorf("CalculateBlastRadius failed: %v", err)
	}
}

func TestDestructionEngineEstimateRecoveryTime(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.EstimateRecoveryTime()
	if err != nil {
		t.Errorf("EstimateRecoveryTime failed: %v", err)
	}
}

func TestDestructionEngineRun(t *testing.T) {
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
