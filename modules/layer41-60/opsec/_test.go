package opsec

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineCommChannelSetup(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.CommChannelSetup()
	if err != nil {
		t.Errorf("CommChannelSetup failed: %v", err)
	}
}

func TestEngineanalyzeChannel(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.analyzeChannel()
	if err != nil {
		t.Errorf("analyzeChannel failed: %v", err)
	}
}

func TestEnginecalculateChannelRisk(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.calculateChannelRisk()
	if err != nil {
		t.Errorf("calculateChannelRisk failed: %v", err)
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
