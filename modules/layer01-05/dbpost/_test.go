package dbpost

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestDBPostEngineExploit(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Exploit()
	if err != nil {
		t.Errorf("Exploit failed: %v", err)
	}
}

func TestDBPostEngineExploitWithTechnique(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ExploitWithTechnique()
	if err != nil {
		t.Errorf("ExploitWithTechnique failed: %v", err)
	}
}

func TestDBPostEngineexploitOracle(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.exploitOracle()
	if err != nil {
		t.Errorf("exploitOracle failed: %v", err)
	}
}

func TestDBPostEngineRun(t *testing.T) {
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
