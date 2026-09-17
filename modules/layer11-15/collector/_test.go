package collector

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestCollectorEngineCollect(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Collect()
	if err != nil {
		t.Errorf("Collect failed: %v", err)
	}
}

func TestCollectorEnginecollectAll(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.collectAll()
	if err != nil {
		t.Errorf("collectAll failed: %v", err)
	}
}

func TestCollectorEngineCaptureScreen(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.CaptureScreen()
	if err != nil {
		t.Errorf("CaptureScreen failed: %v", err)
	}
}

func TestCollectorEngineRun(t *testing.T) {
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
