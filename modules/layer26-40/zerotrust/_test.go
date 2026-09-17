package zerotrust

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineMFABypass(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.MFABypass()
	if err != nil {
		t.Errorf("MFABypass failed: %v", err)
	}
}

func TestEngineSSOAbuse(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.SSOAbuse()
	if err != nil {
		t.Errorf("SSOAbuse failed: %v", err)
	}
}

func TestEngineConditionalAccessBypass(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ConditionalAccessBypass()
	if err != nil {
		t.Errorf("ConditionalAccessBypass failed: %v", err)
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
