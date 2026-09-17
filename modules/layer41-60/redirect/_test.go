package redirect

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineParamManipulation(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ParamManipulation()
	if err != nil {
		t.Errorf("ParamManipulation failed: %v", err)
	}
}

func TestEnginegenerateParamManipulationPayloads(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.generateParamManipulationPayloads()
	if err != nil {
		t.Errorf("generateParamManipulationPayloads failed: %v", err)
	}
}

func TestEngineformatPayloads(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.formatPayloads()
	if err != nil {
		t.Errorf("formatPayloads failed: %v", err)
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
