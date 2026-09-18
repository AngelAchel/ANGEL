package cachesmuggle

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineRequestSmuggle(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.RequestSmuggle()
	if err != nil {
		t.Errorf("RequestSmuggle failed: %v", err)
	}
}

func TestEnginegenerateSmuggleVariants(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.generateSmuggleVariants()
	if err != nil {
		t.Errorf("generateSmuggleVariants failed: %v", err)
	}
}

func TestEngineformatVariants(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.formatVariants()
	if err != nil {
		t.Errorf("formatVariants failed: %v", err)
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
