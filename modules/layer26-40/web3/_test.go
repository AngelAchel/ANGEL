package web3

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineReentrancyDetect(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ReentrancyDetect()
	if err != nil {
		t.Errorf("ReentrancyDetect failed: %v", err)
	}
}

func TestEngineFlashLoanAttack(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.FlashLoanAttack()
	if err != nil {
		t.Errorf("FlashLoanAttack failed: %v", err)
	}
}

func TestEngineOracleManipulation(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.OracleManipulation()
	if err != nil {
		t.Errorf("OracleManipulation failed: %v", err)
	}
}

func TestEngineRun(t *testing.T) {
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
