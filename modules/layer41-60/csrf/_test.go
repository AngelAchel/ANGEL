package csrf

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineTokenBypass(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.TokenBypass()
	if err != nil {
		t.Errorf("TokenBypass failed: %v", err)
	}
}

func TestEngineestimateEntropy(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.estimateEntropy()
	if err != nil {
		t.Errorf("estimateEntropy failed: %v", err)
	}
}

func TestEnginedetectPattern(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.detectPattern()
	if err != nil {
		t.Errorf("detectPattern failed: %v", err)
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
