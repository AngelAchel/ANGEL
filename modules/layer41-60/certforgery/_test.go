package certforgery

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineSelfSignForge(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.SelfSignForge()
	if err != nil {
		t.Errorf("SelfSignForge failed: %v", err)
	}
}

func TestEnginegenerateSerial(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.generateSerial()
	if err != nil {
		t.Errorf("generateSerial failed: %v", err)
	}
}

func TestEngineLetEncryptAbuse(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.LetEncryptAbuse()
	if err != nil {
		t.Errorf("LetEncryptAbuse failed: %v", err)
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
