package destructionchain

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestDestructionChainEngineFullScopeAttack(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.FullScopeAttack()
	if err != nil {
		t.Errorf("FullScopeAttack failed: %v", err)
	}
}

func TestDestructionChainEngineExecuteChain(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ExecuteChain()
	if err != nil {
		t.Errorf("ExecuteChain failed: %v", err)
	}
}

func TestDestructionChainEngineCalculateImpact(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.CalculateImpact()
	if err != nil {
		t.Errorf("CalculateImpact failed: %v", err)
	}
}

func TestDestructionChainEngineRun(t *testing.T) {
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
