package c2server

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineName(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Name()
	if err != nil {
		t.Errorf("Name failed: %v", err)
	}
}

func TestEngineTimestamp(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Timestamp()
	if err != nil {
		t.Errorf("Timestamp failed: %v", err)
	}
}

func TestEngineExecuteCommand(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ExecuteCommand()
	if err != nil {
		t.Errorf("ExecuteCommand failed: %v", err)
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
