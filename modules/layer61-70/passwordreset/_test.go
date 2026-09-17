package passwordreset

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineTokenPredictable(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.TokenPredictable()
	if err != nil {
		t.Errorf("TokenPredictable failed: %v", err)
	}
}

func TestEngineHostHeaderInject(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.HostHeaderInject()
	if err != nil {
		t.Errorf("HostHeaderInject failed: %v", err)
	}
}

func TestEngineResetTokenLeak(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ResetTokenLeak()
	if err != nil {
		t.Errorf("ResetTokenLeak failed: %v", err)
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
