package tls13

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineDowngradeAttack(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.DowngradeAttack()
	if err != nil {
		t.Errorf("DowngradeAttack failed: %v", err)
	}
}

func TestEngineanalyzeDowngrade(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.analyzeDowngrade()
	if err != nil {
		t.Errorf("analyzeDowngrade failed: %v", err)
	}
}

func TestEnginePaddingOracle(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.PaddingOracle()
	if err != nil {
		t.Errorf("PaddingOracle failed: %v", err)
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
