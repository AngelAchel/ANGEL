package upload

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineExtensionBypass(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ExtensionBypass()
	if err != nil {
		t.Errorf("ExtensionBypass failed: %v", err)
	}
}

func TestEnginegenerateExtensionBypasses(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.generateExtensionBypasses()
	if err != nil {
		t.Errorf("generateExtensionBypasses failed: %v", err)
	}
}

func TestEngineformatBypasses(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.formatBypasses()
	if err != nil {
		t.Errorf("formatBypasses failed: %v", err)
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
