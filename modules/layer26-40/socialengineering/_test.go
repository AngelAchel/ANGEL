package socialengineering

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineEmailPhish(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.EmailPhish()
	if err != nil {
		t.Errorf("EmailPhish failed: %v", err)
	}
}

func TestEngineSpearPhish(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.SpearPhish()
	if err != nil {
		t.Errorf("SpearPhish failed: %v", err)
	}
}

func TestEngineVishing(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Vishing()
	if err != nil {
		t.Errorf("Vishing failed: %v", err)
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
