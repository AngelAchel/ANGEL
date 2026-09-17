package racecond

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineTOCTOUExploit(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.TOCTOUExploit()
	if err != nil {
		t.Errorf("TOCTOUExploit failed: %v", err)
	}
}

func TestEngineDoubleFetch(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.DoubleFetch()
	if err != nil {
		t.Errorf("DoubleFetch failed: %v", err)
	}
}

func TestEngineSymlinkRace(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.SymlinkRace()
	if err != nil {
		t.Errorf("SymlinkRace failed: %v", err)
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
