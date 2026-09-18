package ai

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEnginePromptInjection(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.PromptInjection()
	if err != nil {
		t.Errorf("PromptInjection failed: %v", err)
	}
}

func TestEngineModelStealing(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ModelStealing()
	if err != nil {
		t.Errorf("ModelStealing failed: %v", err)
	}
}

func TestEngineAdversarialExample(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.AdversarialExample()
	if err != nil {
		t.Errorf("AdversarialExample failed: %v", err)
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
